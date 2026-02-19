package handler

import "go.uber.org/zap"

type service interface {
	converterService
}

type Handler struct {
	log     *zap.SugaredLogger
	service service
}

func NewHandler(service service, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		log:     logger,
		service: service,
	}
}
