package service

import (
	"context"
	"counter/internal/model"
	"counter/internal/storage"
	"time"
)

type BannerService interface {
	RegisterClick(ctx context.Context, bannerID int) error
	GetStatistics(ctx context.Context, bannerID int, from, to time.Time) ([]*model.BannerClick, error)
}
type bannerService struct {
	repo  storage.BannerStorage
	cache storage.MemoryStorage
}

func NewBannerService(repo storage.BannerStorage, cache storage.MemoryStorage) BannerService {
	return &bannerService{
		repo:  repo,
		cache: cache,
	}
}
