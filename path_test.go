package aliyundrive

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAliyunDrive_Path(t *testing.T) {
	cli := NewDrive(t)
	assert.NotNil(t, cli)

	self, err := cli.MySelf(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, self)
	assert.NotEmpty(t, self.UserID)

	pathInfo, err := cli.Path(context.TODO(), &PathReq{
		DriveID: self.DefaultDriveID,
		FileID:  RootFileID,
	})

	assert.NoError(t, err)
	assert.NotNil(t, pathInfo)
}
