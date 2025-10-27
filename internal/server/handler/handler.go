package handler

import (
	"log"
	"net/http"

	handlerConfig "github.com/ar4ie13/metrics/internal/server/handler/config"
	"github.com/go-chi/chi/v5"
)

type Service interface {
	SaveMetric(metricName string, metricType string, value string) error
	GetAllMetrics() (string, error)
	GetSpecificMetric(metricName string, metricType string) (string, error)
}

type Handler struct {
	service Service
	config  handlerConfig.Config
}

func NewHandler(s Service, c handlerConfig.Config) *Handler {
	return &Handler{s, c}
}

func (h *Handler) ListenAndServe() error {
	router := chi.NewRouter()
	// Uncomment row below when you need to log every request for testing
	//router.Use(middleware.Logger)
	router.Get("/", h.GetAllMetrics)
	router.Get("/value/{metricType}/{metricName}", h.GetMetric)
	router.Post("/update/{metricType}/{metricName}/{metricValue}", h.PostUpdate)

	log.Println("Listening on", h.config.LocalServerAddr)
	if err := http.ListenAndServe(h.config.LocalServerAddr, router); err != nil {
		return err
	}

	return nil
}

func (h *Handler) PostUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	metricName := chi.URLParam(r, "metricName")
	metricType := chi.URLParam(r, "metricType")
	metricValue := chi.URLParam(r, "metricValue")
	err := h.service.SaveMetric(metricName, metricType, metricValue)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}
}

func (h *Handler) GetAllMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	result, _ := h.service.GetAllMetrics()
	_, err := w.Write([]byte(result))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	metricName := chi.URLParam(r, "metricName")
	metricType := chi.URLParam(r, "metricType")
	metric, err := h.service.GetSpecificMetric(metricName, metricType)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_, err = w.Write([]byte(metric))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
