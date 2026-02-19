package router

import (
	"fmt"
	"go-unit-converter/internal/config"
	"go-unit-converter/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerInterface interface {
	converter
}

type converter interface {
	ConvertLengthHandler(c *gin.Context)
	ConvertWeightHandler(c *gin.Context)
	ConvertTemperatureHandler(c *gin.Context)
}

func NewServer(cfg *config.Config, handler handlerInterface) (*http.Server, error) {
	var r *gin.Engine
	switch cfg.Environment {
	case "dev":
		r = gin.Default()
	case "prod":
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Recovery(), gin.Logger())
	default:
		return nil, fmt.Errorf("Не верно указано окружение %s", cfg.Environment)
	}
	r.Use(middleware.CORSMiddleware(cfg.CORS))

	api := r.Group("/api/convert")
	api.POST("/length", handler.ConvertLengthHandler)
	api.POST("/weight", handler.ConvertWeightHandler)
	api.POST("/temperature", handler.ConvertTemperatureHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return server, nil
}
