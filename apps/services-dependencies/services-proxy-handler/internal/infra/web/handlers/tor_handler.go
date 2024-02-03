package handlers

import (
	"apps/services-dependencies/services-proxy-handler/internal/usecase"
	"encoding/json"
	"net/http"
)

type TorProxyHandler struct {
}

func NewTorProxyHandler() *TorProxyHandler {
	return &TorProxyHandler{}
}

func (h *TorProxyHandler) GetTorIPRotation(w http.ResponseWriter, r *http.Request) {
	getTorIPRotation := usecase.NewGetTorProxyRotateUseCase()
	ip, err := getTorIPRotation.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(ip)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
