package aliyundrive

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAliyunDrive_Get(t *testing.T) {
	cli := NewDrive(t)
	assert.NotNil(t, cli)

	self, err := cli.MySelf(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, self)
	assert.NotEmpty(t, self.UserID)

	result, err := cli.UploadFile(context.TODO(), &UploadFileReq{
		DriveID:       self.DefaultDriveID,
		ParentID:      RootFileID,
		FilePath:      "/etc/hosts",
		CheckNameMode: ModeAutoRename,
	})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.FileID)

	info, err := cli.Get(context.TODO(), &GetFileReq{
		DriveID: self.DefaultDriveID,
		FileID:  result.FileID,
	})
	assert.NoError(t, err)
	assert.NotNil(t, info)
	assert.NotEmpty(t, info.FileID)

	err = cli.Trash(context.TODO(), &DeleteFileReq{
		DriveID: self.DefaultDriveID,
		FileID:  result.FileID,
	})
	assert.NoError(t, err)
}
