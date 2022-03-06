package aliyundrive

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAliyunDrive_Move(t *testing.T) {
	cli := NewDrive(t)
	assert.NotNil(t, cli)

	self, err := cli.MySelf(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, self)
	assert.NotEmpty(t, self.UserID)

	file, err := cli.UploadFile(context.TODO(), &UploadFileReq{
		DriveID:       self.DefaultDriveID,
		ParentID:      RootFileID,
		FilePath:      "/etc/hosts",
		CheckNameMode: ModeAutoRename,
	})

	assert.NoError(t, err)
	assert.NotNil(t, file)
	assert.NotEmpty(t, file.FileID)

	folder, err := cli.CreateFolder(context.TODO(), &CreateFolderReq{
		DriveID:       self.DefaultDriveID,
		ParentFileID:  RootFileID,
		Name:          "moved",
		CheckNameMode: ModeRefuse,
	})

	assert.NoError(t, err)
	assert.NotNil(t, folder)

	moveResult, err := cli.Move(context.TODO(), &MoveReq{
		DriveID:        self.DefaultDriveID,
		FileID:         file.FileID,
		ToParentFileID: folder.FileID,
	})
	assert.NoError(t, err)
	assert.NotNil(t, moveResult)
	assert.NotEmpty(t, moveResult.FileID)

	err = cli.Trash(context.TODO(), &DeleteFileReq{
		DriveID: self.DefaultDriveID,
		FileID:  folder.FileID,
	})
	assert.NoError(t, err)
}
