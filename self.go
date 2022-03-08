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

type GetSelfUserResp struct {
	DomainID                    string      `json:"domain_id"`
	UserID                      string      `json:"user_id"`
	Avatar                      string      `json:"avatar"`
	CreatedAt                   int64       `json:"created_at"`
	UpdatedAt                   int64       `json:"updated_at"`
	Email                       string      `json:"email"`
	NickName                    string      `json:"nick_name"`
	Phone                       string      `json:"phone"`
	Role                        string      `json:"role"`
	Status                      string      `json:"status"`
	UserName                    string      `json:"user_name"`
	Description                 string      `json:"description"`
	DefaultDriveID              string      `json:"default_drive_id"`
	DenyChangePasswordBySelf    bool        `json:"deny_change_password_by_self"`
	NeedChangePasswordNextLogin bool        `json:"need_change_password_next_login"`
	Permission                  interface{} `json:"permission"`
}

func (r *AliyunDrive) MySelf(ctx context.Context) (*GetSelfUserResp, error) {
	var result GetSelfUserResp

	req := config{
		Method: http.MethodPost,
		URL:    "https://api.aliyundrive.com/v2/user/get",
		Body:   "{}",
	}

	_, err := r.request(ctx, &req, &result)
	if err != nil {
		return nil, err
	}

	if result.UserID == "" {
		return nil, fmt.Errorf("user id is empty, pls check")
	}

	return &result, nil
}
