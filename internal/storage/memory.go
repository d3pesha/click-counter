package storage

import (
	"sync"
)

type memoryStorage struct {
	mu     sync.Mutex
	clicks map[int]int
}

func NewMemoryStorage() *memoryStorage {
	return &memoryStorage{
		clicks: make(map[int]int),
	}
}

func (s *memoryStorage) IncrementClick(bannerID int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.clicks[bannerID]++
}

func (s *memoryStorage) GetAndClearClicks() map[int]int {
	s.mu.Lock()
	defer s.mu.Unlock()

	data := s.clicks
	s.clicks = make(map[int]int)

	return data
}

func (s *memoryStorage) GetByKey(key int) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	data, ok := s.clicks[key]
	if ok {
		delete(s.clicks, key)
	}

	return data
}
