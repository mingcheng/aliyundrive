package aliyundrive

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAliyunDrive_Rename(t *testing.T) {
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

	renameResult, err := cli.Rename(context.TODO(), &RenameFileReq{
		DriveID: self.DefaultDriveID,
		FileID:  result.FileID,
		Name:    "hosts.txt",
	})

	assert.NoError(t, err)
	assert.NotNil(t, renameResult)

	err = cli.Trash(context.TODO(), &DeleteFileReq{
		DriveID: self.DefaultDriveID,
		FileID:  result.FileID,
	})
	assert.NoError(t, err)
}
