package aliyundrive

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAliyunDrive_Token(t *testing.T) {
	cli := NewDrive(t)

	myself, err := cli.MySelf(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, myself)

	// TODO
}
