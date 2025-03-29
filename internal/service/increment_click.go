package service

import (
	"context"
	"log"
)

func (s *bannerService) RegisterClick(ctx context.Context, bannerID int) error {
	if _, exist := s.cache.GetByKey(bannerID); exist {
		s.cache.Increment(bannerID)
		log.Printf("Click registered for bannerID: %d", bannerID)

		return nil
	}

	_, err := s.bannerStorage.GetByID(ctx, bannerID)
	if err != nil {
		return err
	}

	s.cache.SetNewBanner(bannerID, 1)
	log.Printf("Click registered for bannerID: %d", bannerID)

	return nil
}
