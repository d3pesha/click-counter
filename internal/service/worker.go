package service

import (
	"context"
	"counter/internal/storage"
	"log"
	"time"
)

type Worker struct {
	clickStorage storage.BannerClickStorage
	cache        storage.MemoryStorage
}

func NewWorker(clickStorage storage.BannerClickStorage, cache storage.MemoryStorage) *Worker {
	return &Worker{
		clickStorage: clickStorage,
		cache:        cache,
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
	clicks := w.cache.GetAndClear()
	now := time.Now().Truncate(time.Minute)

	for bannerID, count := range clicks {
		if count <= 0 {
			continue
		}

		err := w.clickStorage.IncrementClick(ctx, bannerID, now, count)
		if err != nil {
			log.Printf("Error saving clicks for bannerID %d: %v\n", bannerID, err)
		} else {
			log.Printf("Saved %d clicks for bannerID %d\n", count, bannerID)
		}
	}
}
