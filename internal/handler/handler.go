package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

type Service interface {
	SaveMetric(string, string, string) error
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
	router.Use(middleware.Logger)
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
