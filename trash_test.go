package aliyundrive

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAliyunDrive_Trash(t *testing.T) {
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

	err = cli.Trash(context.TODO(), &DeleteFileReq{
		DriveID: self.DefaultDriveID,
		FileID:  result.FileID,
	})
	assert.NoError(t, err)
}

func TestAliyunDrive_TrashFolders(t *testing.T) {
	cli := NewDrive(t)
	assert.NotNil(t, cli)

	self, err := cli.MySelf(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, self)
	assert.NotEmpty(t, self.UserID)

	folderResp, err := cli.CreateFolder(context.TODO(), &CreateFolderReq{
		DriveID:       self.DefaultDriveID,
		ParentFileID:  RootFileID,
		Name:          "test",
		CheckNameMode: ModeRefuse,
	})

	assert.NoError(t, err)
	assert.NotNil(t, folderResp)
	assert.NotEmpty(t, folderResp.FileID)

	err = cli.Trash(context.TODO(), &DeleteFileReq{
		DriveID: self.DefaultDriveID,
		FileID:  folderResp.FileID,
	})

	assert.NoError(t, err)
}
