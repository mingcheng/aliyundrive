package store

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewFileStore(t *testing.T) {
	fileStorePath := os.Getenv("FILE_STORE_PATH")
	if fileStorePath == "" {
		return
	}

	store, err := NewFileStore(fileStorePath)

	assert.NoError(t, err)
	assert.NotNil(t, store)

	got, err := store.Get(context.TODO(), "not-exists")
	assert.Error(t, err)
	assert.Nil(t, got)

	err = store.Set(context.TODO(), "hello", []byte("world"))
	assert.NoError(t, err)

	get, err := store.Get(context.TODO(), "hello")
	assert.NoError(t, err)
	assert.Equal(t, get, []byte("world"))
}
