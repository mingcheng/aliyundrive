/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package aliyundrive

import (
	"context"
	"fmt"
	"net/http"
)

type DeleteFileReq struct {
	DriveID string `json:"drive_id"`
	FileID  string `json:"file_id"`
}

type DeleteFileResp struct {
	DomainID    string `json:"domain_id"`
	DriveID     string `json:"drive_id"`
	FileID      string `json:"file_id"`
	AsyncTaskID string `json:"async_task_id"`
}

func (r *AliyunDrive) Trash(ctx context.Context, request *DeleteFileReq) error {
	response, err := r.request(ctx, &config{
		Method: http.MethodPost,
		URL:    "https://api.aliyundrive.com/v2/recyclebin/trash",
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
