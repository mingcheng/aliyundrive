package aliyundrive

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAliyunDrive_CreateFolder(t *testing.T) {
	cli := NewDrive(t)

	self, err := cli.MySelf(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, self)
	assert.NotEmpty(t, self.UserID)

	result, err := cli.CreateFolder(context.TODO(), &CreateFolderReq{
		DriveID:      self.DefaultDriveID,
		ParentFileID: RootFileID,
		Name:         "just-for-example",
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

func TestAliyunDrive_CreateFolder2(t *testing.T) {
	cli := NewDrive(t)

	self, err := cli.MySelf(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, self)
	assert.NotEmpty(t, self.UserID)

	for i := 0; i < 10; i++ {
		result, err := cli.CreateFolder(context.TODO(), &CreateFolderReq{
			DriveID:       self.DefaultDriveID,
			ParentFileID:  RootFileID,
			Name:          "just/for/example",
			CheckNameMode: ModeRefuse,
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.FileID)
		assert.Equal(t, "example", result.FileName)

		err = cli.Trash(context.TODO(), &DeleteFileReq{
			DriveID: self.DefaultDriveID,
			FileID:  result.FileID,
		})
		assert.NoError(t, err)
	}
}
