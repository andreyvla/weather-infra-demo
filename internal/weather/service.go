package weather

import (
	"context"
)

type Service struct {
	client *Client
	cache  *Cache
}

func NewService(client *Client, cache *Cache) *Service {
	return &Service{
		client: client,
		cache:  cache,
	}
}

func (s *Service) Get(ctx context.Context) (*Weather, error) {
	if w, ok := s.cache.Get(); ok {
		return w, nil
	}

	w, err := s.client.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	s.cache.Set(w)
	return w, nil
}
