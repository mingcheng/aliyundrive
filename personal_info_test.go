package aliyundrive

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAliyunDrive_PersonalInfo(t *testing.T) {
	cli := NewDrive(t)

	result, err := cli.PersonalInfo(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, result)

	assert.NotEmpty(t, result.RightsInfo)
	assert.NotEmpty(t, result.SpaceInfo)
}
