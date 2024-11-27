package handler

import (
	"github.com/atadzan/music-library-api/internal/pkg/services"
)

type Handler struct {
	service services.Service
}

func New(service services.Service) Handler {
	return Handler{
		service: service,
	}
}
