package aliyundrive

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAliyunDrive_Lists(t *testing.T) {
	cli := NewDrive(t)
	assert.NotNil(t, cli)

	self, err := cli.MySelf(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, self)
	assert.NotEmpty(t, self.UserID)

	files, err := cli.Lists(context.TODO(), &FileListReq{
		DriveID: self.DefaultDriveID,
	})

	assert.NoError(t, err)
	assert.NotNil(t, files)
}
