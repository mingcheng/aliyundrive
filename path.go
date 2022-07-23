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
	"net/http"
)

type PathReq struct {
	DriveID string `json:"drive_id"`
	FileID  string `json:"file_id"`
}

type PathResp struct {
	Items []*File `json:"items"`
}

func (r *AliyunDrive) Path(ctx context.Context, request *PathReq) (*PathResp, error) {
	var resp PathResp

	_, err := r.request(ctx, &config{
		Method: http.MethodPost,
		URL:    "https://api.aliyundrive.com/adrive/v1/file/get_path",
		Body:   request,
	}, &resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}
