package controller

import (
	"github.com/MorZLE/ParseTSVBiocad/internal/service"
)

func NewHandler(s service.Service) *Handler {
	return &Handler{s}
}

type Handler struct {
	s service.Service
}

func (h *Handler) Start() {
}
