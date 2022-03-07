package aliyundrive

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAliyunDrive_UploadFile(t *testing.T) {
	cli := NewDrive(t)
	assert.NotNil(t, cli)

	self, err := cli.MySelf(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, self)
	assert.NotEmpty(t, self.UserID)

	for i := 0; i < 10; i++ {
		result, err := cli.UploadFile(context.TODO(), &UploadFileReq{
			DriveID:       self.DefaultDriveID,
			ParentID:      RootFileID,
			FilePath:      "/etc/hosts",
			CheckNameMode: ModeAutoRename,
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.FileID)
		assert.True(t, result.Size > 0)

		err = cli.Trash(context.TODO(), &DeleteFileReq{
			DriveID: self.DefaultDriveID,
			FileID:  result.FileID,
		})
		assert.NoError(t, err)
	}
}
