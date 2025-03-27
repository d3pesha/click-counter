package service

import (
	"context"
	"log"
)

func (s *bannerService) RegisterClick(ctx context.Context, bannerID int) error {
	_, err := s.repo.GetByID(ctx, bannerID)
	if err != nil {
		return err
	}

	s.cache.IncrementClick(bannerID)
	log.Printf("Click registered for bannerID: %d", bannerID)

	return nil
}
