package health

import (
	"encoding/json"
	"net/http"
	"time"
)

// структура ответа health-check
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
	Version   string    `json:"version"`
}

// структура обработчика
type Handler struct {
	service string
	version string
}

// конструктор для создания обработчика
func NewHandler(service, version string) *Handler {
	return &Handler{
		service: service,
		version: version,
	}
}

//  делает Handler совместимым с http.Handler
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now(),
		Service:   h.service,
		Version:   h.version,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}