package service

import (
	"context"
	"counter/internal/model"
	"log"
	"time"
)

func (s *bannerService) GetStatistics(ctx context.Context, bannerID int, from, to time.Time) ([]*model.BannerClick, error) {
	if err := s.saveCurrentMinuteClicks(ctx, bannerID); err != nil {
		return nil, err
	}

	result, err := s.repo.GetStats(ctx, bannerID, from, to)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *bannerService) saveCurrentMinuteClicks(ctx context.Context, bannerID int) error {
	count := s.cache.GetByKey(bannerID)
	if count == 0 {
		return nil
	}

	now := time.Now().Truncate(time.Minute)

	err := s.repo.IncrementClick(ctx, bannerID, now, count)
	if err != nil {
		log.Printf("Error saving count for bannerID %d at %v: %v", bannerID, now, err)

		return err
	}

	log.Printf("Saved %d clicks for bannerID %d\n", count, bannerID)

	return nil
}
