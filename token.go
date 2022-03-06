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
	"time"
)

type Token struct {
	AccessToken  string    `json:"access_token"`
	ExpiredAt    time.Time `json:"expired_at"` // access-token 的过期时间，秒级
	RefreshToken string    `json:"refresh_token"`
}

type getTokenReq struct {
	Code      string `json:"code"`
	LoginType string `json:"loginType"`
	DeviceId  string `json:"deviceId"`
}

func (r *AliyunDrive) Token(ctx context.Context, request *getTokenReq) (*RefreshTokenResp, error) {
	response, err := r.Client.R().SetBody(getTokenReq{
		Code:      request.Code,
		LoginType: "normal",
		DeviceId:  "aliyundrive",
	}).SetResult(RefreshTokenResp{}).Post("https://api.aliyundrive.com/token/get")

	if err != nil {
		return nil, err
	}

	return response.Result().(*RefreshTokenResp), nil
}
