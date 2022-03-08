package store

import "context"

type MemoryStore struct {
	data map[string][]byte
}

func (s *MemoryStore) Set(ctx context.Context, key string, data []byte) error {
	s.data[key] = data
	return nil
}

func (s *MemoryStore) Get(ctx context.Context, key string) ([]byte, error) {
	return s.data[key], nil
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string][]byte),
	}
}
