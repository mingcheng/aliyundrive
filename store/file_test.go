package store

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFileStore(t *testing.T) {
	store, err := NewFileStore("/Users/mingcheng/.aliyundrive")

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
