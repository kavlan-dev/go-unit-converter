package handler

import (
	"go-unit-converter/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type convertService interface {
	ConvertLength(model.ConversionRequest) (model.ConversionResponse, error)
	ConvertWeight(model.ConversionRequest) (model.ConversionResponse, error)
	ConvertTemperature(model.ConversionRequest) (model.ConversionResponse, error)
}

type convertHandler struct {
	log     *zap.SugaredLogger
	service convertService
}

func NewConvertHandler(service convertService, logger *zap.SugaredLogger) *convertHandler {
	return &convertHandler{
		log:     logger,
		service: service,
	}
}

func (h *convertHandler) ConvertLengthHandler(c *gin.Context) {
	var req model.ConversionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ошибка в запросе"})
		return
	}

	result, err := h.service.ConvertLength(req)
	if err != nil {
		h.log.Errorf("ConvertLength: %w", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при конвертации длины"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *convertHandler) ConvertWeightHandler(c *gin.Context) {
	var req model.ConversionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ошибка в запросе"})
		return
	}

	result, err := h.service.ConvertWeight(req)
	if err != nil {
		h.log.Errorf("ConvertWeight: %w", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при конвертации веса"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *convertHandler) ConvertTemperatureHandler(c *gin.Context) {
	var req model.ConversionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ошибка в запросе"})
		return
	}

	result, err := h.service.ConvertTemperature(req)
	if err != nil {
		h.log.Errorf("ConvertTemperature: %w", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при конвертации температуры"})
		return
	}

	c.JSON(http.StatusOK, result)
}
