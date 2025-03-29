package storage

import (
	"context"
	"counter/internal/model"
	"time"
)

type BannerStorage interface {
	GetByID(ctx context.Context, bannerID int) (*model.Banner, error)
	GetAllIDs(ctx context.Context) ([]int, error)
}

type BannerClickStorage interface {
	IncrementClick(ctx context.Context, bannerID int, timestamp time.Time, countClick int) error
	GetStats(ctx context.Context, bannerID int, tsFrom, tsTo time.Time) ([]*model.BannerClick, error)
}

type MemoryStorage interface {
	Increment(bannerID int)
	GetAndClear() map[int]int
	GetByKey(key int) (int, bool)
	GetByKeyAndClear(key int) (int, bool)
	SetNewBanner(key, value int)
	FillCache(ctx context.Context)
}
