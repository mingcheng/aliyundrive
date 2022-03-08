package aliyundrive

import (
	"context"
	"fmt"
	"net/http"
)

type RestoreFileReq struct {
	DriveId string `json:"drive_id"`
	FileId  string `json:"file_id"`
}

func (r *AliyunDrive) Restore(ctx context.Context, request *RestoreFileReq) error {
	response, err := r.request(ctx, &config{
		Method: http.MethodPost,
		URL:    "https://api.aliyundrive.com/v2/recyclebin/restore",
		Body:   request,
	}, nil)

	if err != nil {
		return err
	}

	if !response.IsSuccess() {
		return fmt.Errorf("%s", response.Status())
	}

	return nil
}
