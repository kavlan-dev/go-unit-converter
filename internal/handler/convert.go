package handler

import (
	"go-unit-converter/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type convertService interface {
	ConvertLength(*model.ConversionRequest) (*model.ConversionResponse, error)
	ConvertWeight(*model.ConversionRequest) (*model.ConversionResponse, error)
	ConvertTemperature(*model.ConversionRequest) (*model.ConversionResponse, error)
}

type ConvertHandler struct {
	log     *zap.SugaredLogger
	service convertService
}

func NewConvertHandler(service convertService, logger *zap.SugaredLogger) *ConvertHandler {
	return &ConvertHandler{
		log:     logger,
		service: service,
	}
}

func (h *ConvertHandler) ConvertLengthHandler(c *gin.Context) {
	var req model.ConversionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка в запросе"})
		return
	}

	result, err := h.service.ConvertLength(&req)
	if err != nil {
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при конвертации длины"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *ConvertHandler) ConvertWeightHandler(c *gin.Context) {
	var req model.ConversionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка в запросе"})
		return
	}

	result, err := h.service.ConvertWeight(&req)
	if err != nil {
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при конвертации веса"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *ConvertHandler) ConvertTemperatureHandler(c *gin.Context) {
	var req model.ConversionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка в запросе"})
		return
	}

	result, err := h.service.ConvertTemperature(&req)
	if err != nil {
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при конвертации температуры"})
		return
	}

	c.JSON(http.StatusOK, result)
}
