package main

import (
	"context"
	"counter/config"
	"counter/internal/api/handler"
	"counter/internal/database"
	"counter/internal/service"
	"counter/internal/storage"
	"counter/seed"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.LoadConfig()
	db := database.NewPostgresDB(cfg)
	defer db.DB.Close()

	db.Migrate(cfg)
	seed.Banners(db.DB)

	clickStorage := storage.NewBannerClickStorage(db.DB)
	bannerStorage := storage.NewBannerStorage(db.DB)

	memStorage := storage.NewMemoryStorage(bannerStorage)
	memStorage.FillCache(ctx)

	bannerService := service.NewBannerService(bannerStorage, clickStorage, memStorage)

	worker := service.NewWorker(clickStorage, memStorage)
	go worker.Start(ctx)

	app := fiber.New()
	handler.Register(app, bannerService)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := app.Listen(":" + cfg.AppPort); err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}()

	<-sigCh
	log.Println("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	log.Println("Server gracefully stopped")
}
