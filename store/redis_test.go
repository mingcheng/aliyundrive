package store

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewRedisStore(t *testing.T) {
	redisAddr := os.Getenv("REDIS_ADDR")
	assert.NotEmpty(t, redisAddr)

	const name = "test"

	client, err := NewRedisStore(&redis.Options{
		Addr: redisAddr,
	})

	assert.NoError(t, err)
	assert.NotNil(t, client)

	err = client.Set(context.TODO(), name, []byte("OK"))
	assert.NoError(t, err)

	data, err := client.Get(context.TODO(), name)
	assert.NoError(t, err)
	assert.Equal(t, "OK", string(data))
}
