package store

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"syscall"
)

type FileStore struct {
	Path string
}

func (f FileStore) Get(ctx context.Context, key string) ([]byte, error) {
	getPath, err := f.getPath(key)
	if err != nil {
		return nil, err
	}

	bs, err := ioutil.ReadFile(getPath)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (f FileStore) Set(ctx context.Context, key string, data []byte) error {
	getPath, err := f.getPath(key)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(getPath, data, 0o600)
}

func (f FileStore) getPath(key string) (string, error) {
	return path.Join(f.Path, key+".txt"), nil
}

func NewFileStore(dir string) (*FileStore, error) {
	info, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("%s must be a directory", info.Name())
	}

	err = syscall.Access(dir, syscall.O_RDWR)
	if err != nil {
		return nil, fmt.Errorf("%s is not writeable", dir)
	}

	return &FileStore{
		Path: dir,
	}, nil
}
