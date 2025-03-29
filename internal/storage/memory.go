package storage

import (
	"context"
	"log"
	"sync"
)

type memoryStorage struct {
	bannerStorage BannerStorage
	mu            sync.Mutex
	clicks        map[int]int
}

func NewMemoryStorage(storage BannerStorage) MemoryStorage {
	return &memoryStorage{
		bannerStorage: storage,
		clicks:        make(map[int]int),
	}
}

func (s *memoryStorage) Increment(bannerID int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.clicks[bannerID]++
}

func (s *memoryStorage) GetByKey(key int) (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, ok := s.clicks[key]

	return data, ok
}

func (s *memoryStorage) GetAndClear() map[int]int {
	s.mu.Lock()
	defer s.mu.Unlock()

	data := make(map[int]int, len(s.clicks))
	for k, v := range s.clicks {
		data[k] = v
		s.clicks[k] = 0
	}

	return data
}

func (s *memoryStorage) GetByKeyAndClear(key int) (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, ok := s.clicks[key]
	if !ok {
		return 0, false
	}

	s.clicks[key] = 0

	return data, true
}

func (s *memoryStorage) SetNewBanner(key, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.clicks[key] = value
}

func (s *memoryStorage) FillCache(ctx context.Context) {
	s.mu.Lock()
	defer s.mu.Unlock()

	bannerIDs, err := s.bannerStorage.GetAllIDs(ctx)
	if err != nil {
		log.Printf("Error getting banner IDs: %v\n", err)
	}

	for _, bannerID := range bannerIDs {
		s.clicks[bannerID] = 0
	}

	log.Println("cache filled successfully")
}
