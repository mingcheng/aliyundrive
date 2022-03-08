package aliyundrive

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAliyunDrive_CreateShare(t *testing.T) {
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

	shareFile, err := cli.CreateShare(context.TODO(), &CreateShareReq{
		DriveId:    self.DefaultDriveID,
		FileIdList: []string{result.FileID},
	})
	assert.NoError(t, err)
	assert.NotNil(t, shareFile)
	assert.NotEmpty(t, shareFile.FileId)

	shareUpdateResult, err := cli.UpdateShare(context.TODO(), &UpdateShareReq{
		ShareId:    shareFile.ShareId,
		Expiration: "",
		//expiration: "2022-04-07T03:00:59.588Z"
	})
	assert.NoError(t, err)
	assert.NotNil(t, shareUpdateResult)
	assert.NotEmpty(t, shareUpdateResult.ShareId)

	err = cli.CancelShare(context.TODO(), &CancelShareReq{
		ShareId: shareFile.ShareId,
	})
	assert.NoError(t, err)

	err = cli.Trash(context.TODO(), &DeleteFileReq{
		DriveID: self.DefaultDriveID,
		FileID:  result.FileID,
	})
	assert.NoError(t, err)
}

func TestAliyunDrive_ListShare(t *testing.T) {
	cli := NewDrive(t)
	assert.NotNil(t, cli)

	self, err := cli.MySelf(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, self)
	assert.NotEmpty(t, self.UserID)

	sharedFiles, err := cli.ListShare(context.TODO(), &ListShareReq{
		Creator:         self.UserID,
		IncludeCanceled: false,
		Category:        TypeFile,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, sharedFiles.Items)
}

func TestAliyunDrive_UpdateShare(t *testing.T) {

}
