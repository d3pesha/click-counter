package service

import (
	"context"
	"counter/internal/model"
	"counter/internal/storage"
	"time"
)

type bannerService struct {
	bannerStorage storage.BannerStorage
	clickStorage  storage.BannerClickStorage
	cache         storage.MemoryStorage
}

func NewBannerService(
	bannerStorage storage.BannerStorage,
	clickStorage storage.BannerClickStorage,
	cache storage.MemoryStorage,
) BannerService {
	return &bannerService{
		bannerStorage: bannerStorage,
		clickStorage:  clickStorage,
		cache:         cache,
	}
}

type BannerService interface {
	RegisterClick(ctx context.Context, bannerID int) error
	GetStatistics(ctx context.Context, bannerID int, from, to time.Time) ([]*model.BannerClick, error)
}
