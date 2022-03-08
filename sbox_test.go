package aliyundrive

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAliyunDrive_GetSBox(t *testing.T) {
	cli := NewDrive(t)
	assert.NotNil(t, cli)

	self, err := cli.MySelf(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, self)
	assert.NotEmpty(t, self.UserID)

	sbox, err := cli.GetSBox(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, sbox)
	assert.True(t, sbox.TotalSize > 0)
}
