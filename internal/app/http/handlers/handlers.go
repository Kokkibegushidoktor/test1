package handlers

import "github.com/Kokkibegushidoktor/test1/internal/service"

type Handler struct {
	services *service.Services
}

func New(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}
