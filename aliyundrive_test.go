package aliyundrive

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/mingcheng/aliyundrive/store"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func NewDrive(t *testing.T) *AliyunDrive {
	redisAddr := os.Getenv("REDIS_ADDR")
	assert.NotEmpty(t, redisAddr)

	client, err := store.NewRedisStore(&redis.Options{
		Addr: redisAddr,
	})

	assert.NoError(t, err)
	assert.NotNil(t, client)

	// using `JSON.parse(localStorage.token).refresh_token` to get refresh token from web
	refreshToken, err := client.Get(context.TODO(), KeyRefreshToken)
	if len(refreshToken) == 0 {
		refreshToken = []byte(os.Getenv("REFRESH_TOKEN"))
		assert.NotEmpty(t, refreshToken)
	}

	cli := New(WithStore(client))

	if !cli.IsLogin(context.TODO()) {
		token, err := cli.RefreshToken(context.TODO(), &RefreshTokenReq{
			RefreshToken: string(refreshToken),
		})

		assert.NoError(t, err)
		assert.NotNil(t, token)
		assert.NotEmpty(t, token.RefreshToken)
	}

	assert.NotNil(t, cli)
	return cli
}
