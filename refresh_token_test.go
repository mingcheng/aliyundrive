package aliyundrive

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthService_RefreshToken(t *testing.T) {
	cli := NewDrive(t)
	assert.NotNil(t, cli)

	resp, err := cli.MySelf(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	refreshToken, err := cli.store.Get(context.TODO(), KeyRefreshToken)
	assert.NoError(t, err)
	assert.NotEmpty(t, refreshToken)

	token, err := cli.RefreshToken(context.TODO(), &RefreshTokenReq{
		RefreshToken: string(refreshToken),
	})
	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.NotEmpty(t, token.RefreshToken)
	assert.NotEmpty(t, token.AccessToken)
}
