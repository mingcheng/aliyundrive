package aliyundrive

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAliyunDrive_DeviceList(t *testing.T) {
	cli := NewDrive(t)

	result, err := cli.DeviceList(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Result)
}
