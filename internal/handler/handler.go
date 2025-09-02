package handler

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Service interface {
	SaveMetric(string, string, string) error
	GetAllMetrics() string
	GetSpecificMetric(string, string) (string, error)
}

type Config interface {
	GetLocalServerAddr() string
}

type Handler struct {
	s Service
	c Config
}

func NewHandler(s Service, c Config) *Handler {
	return &Handler{s, c}
}

func (h *Handler) ListenAndServe() error {
	router := chi.NewRouter()
	// Uncomment row below when you need to log every request for testing
	//router.Use(middleware.Logger)
	router.Get("/", h.GetAllMetrics)
	router.Get("/value/{metricType}/{metricName}", h.GetMetric)
	router.Post("/update/{metricType}/{metricName}/{metricValue}", h.PostUpdate)

	log.Println("Listening on", h.c.GetLocalServerAddr())
	if err := http.ListenAndServe(h.c.GetLocalServerAddr(), router); err != nil {
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
	err := h.s.SaveMetric(metricName, metricType, metricValue)
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

	result := h.s.GetAllMetrics()
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
	metric, err := h.s.GetSpecificMetric(metricName, metricType)
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
