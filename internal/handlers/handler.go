package handlers

import (
	"lumel/internal/services"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger  *logrus.Logger
	service services.Service
}

func NewHandler(logger *logrus.Logger, service services.Service) *Handler {
	return &Handler{logger: logger, service: service}
}
