package service

import (
	"context"
	"counter/internal/storage"
	"log"
	"time"
)

type Worker struct {
	repo  storage.BannerStorage
	cache storage.MemoryStorage
}

func NewWorker(repo storage.BannerStorage, cache storage.MemoryStorage) *Worker {
	return &Worker{
		repo:  repo,
		cache: cache,
	}
}

func (w *Worker) Start(ctx context.Context) {
	now := time.Now()
	nextMinute := now.Add(time.Minute).Truncate(time.Minute)
	waitDuration := nextMinute.Sub(now)
	time.Sleep(waitDuration)

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	log.Println("Worker started")

	w.processClicks(ctx) // saving clicks before start at 00

	for {
		select {
		case <-ctx.Done():
			log.Println("Worker stopped")
			return
		case <-ticker.C:
			w.processClicks(ctx)
		}
	}
}

func (w *Worker) processClicks(ctx context.Context) {
	clicks := w.cache.GetAndClearClicks()
	now := time.Now().Truncate(time.Minute)

	for bannerID, count := range clicks {
		err := w.repo.IncrementClick(ctx, bannerID, now, count)
		if err != nil {
			log.Printf("Error saving clicks for bannerID %d: %v\n", bannerID, err)
		} else {
			log.Printf("Saved %d clicks for bannerID %d\n", count, bannerID)
		}
	}
}
