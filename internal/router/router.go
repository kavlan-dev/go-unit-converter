package router

import (
	"fmt"
	"go-unit-converter/internal/config"
	"go-unit-converter/internal/middleware"

	"github.com/gin-gonic/gin"
)

type convertHandler interface {
	ConvertLengthHandler(c *gin.Context)
	ConvertWeightHandler(c *gin.Context)
	ConvertTemperatureHandler(c *gin.Context)
}

func New(cfg *config.Config, handler convertHandler) (*gin.Engine, error) {
	var r *gin.Engine
	switch cfg.Environment {
	case "dev":
		r = gin.Default()
	case "prod":
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Recovery(), gin.Logger())
	default:
		return nil, fmt.Errorf("не верно указано окружение %s", cfg.Environment)
	}
	r.Use(middleware.CORSMiddleware(cfg.CORS))

	api := r.Group("/api/convert")
	api.POST("/length", handler.ConvertLengthHandler)
	api.POST("/weight", handler.ConvertWeightHandler)
	api.POST("/temperature", handler.ConvertTemperatureHandler)

	return r, nil
}
