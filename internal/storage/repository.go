package storage

import (
	"context"
	"counter/internal/model"
	"time"
)

type BannerStorage interface {
	GetByID(ctx context.Context, bannerID int) (*model.Banner, error)
	IncrementClick(ctx context.Context, bannerID int, timestamp time.Time, countClick int) error
	GetStats(ctx context.Context, bannerID int, tsFrom, tsTo time.Time) ([]*model.BannerClick, error)
}

type MemoryStorage interface {
	IncrementClick(bannerID int)
	GetAndClearClicks() map[int]int
	GetByKey(key int) int
}
