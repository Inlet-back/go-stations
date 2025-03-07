package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{}


// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
//ServeHTTPはHandlerに元からある
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := &model.HealthzResponse{Message: "OK"}

	if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
}
