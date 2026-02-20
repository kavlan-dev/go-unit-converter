package app

import (
	"context"
	"go-unit-converter/internal/config"
	"go-unit-converter/internal/handler"
	"go-unit-converter/internal/router"
	"go-unit-converter/internal/service"
	"go-unit-converter/internal/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Ошибка инициализации конфигурации: %v", err)
	}

	logger, err := util.NewLogger(cfg.Environment)
	if err != nil {
		log.Fatalf("Ошибка инициализации логгера: %v", err)
	}
	defer logger.Sync()

	convertService := service.NewConvertService()
	convertHandler := handler.NewConvertHandler(convertService, logger)

	r, err := router.NewServer(cfg, convertHandler)
	if err != nil {
		logger.Fatalf("Ошибка инициализации сервера: %v", err)
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	logger.Info("Сервер запущен")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		logger.Info("Остановка...")
		cancel()
		if err := server.Shutdown(ctx); err != nil {
			logger.Errorf("Ошибка остановки сервера %v:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
