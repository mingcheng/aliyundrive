package store

import (
	"context"
	"io/ioutil"
)

type FileStore struct {
	Path string
}

func (f FileStore) Get(ctx context.Context, key string) ([]byte, error) {
	bs, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (f FileStore) Set(ctx context.Context, key string, data []byte) error {
	return ioutil.WriteFile(f.Path, data, 0o600)
}

func NewFileStore(file string) *FileStore {
	return &FileStore{
		Path: file,
	}
}
