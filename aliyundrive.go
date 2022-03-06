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

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

const KeyAccessToken = "aliyun_drive_access_token"
const KeyRefreshToken = "aliyun_drive_refresh_token"

const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 Edg/98.0.1108.62"

type AliyunDrive struct {
	logger *logrus.Logger
	client *resty.Client
	store  Store

	accessToken string
}

type config struct {
	Method string
	URL    string
	Body   interface{}
}

func (r *AliyunDrive) request(ctx context.Context, req *config, result interface{}) (*resty.Response, error) {
	request := r.client.R().SetContext(ctx)

	if req.Body != nil {
		request.SetBody(req.Body)
	}

	if result != nil {
		request.SetResult(result)
	}

	response, err := request.Execute(req.Method, req.URL)
	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, fmt.Errorf("%s", response.Status())
	}

	return response, err
}

func (r *AliyunDrive) log(level logrus.Level, args ...interface{}) {
	r.logger.Log(level, args...)
}

func New(options ...OptionFunc) *AliyunDrive {
	client := resty.New()
	r := &AliyunDrive{
		logger: logrus.New(),
		client: client,
	}

	client.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		request.SetHeader("User-Agent", userAgent)
		request.SetHeader("Referer", "https://www.aliyundrive.com/")

		if request.Method == http.MethodPost {
			request.SetHeader("Content-Type", "application/json; charset=utf-8")
		}

		if r.accessToken != "" {
			request.SetAuthToken(r.accessToken)
		}

		return nil
	})

	for _, v := range options {
		if v != nil {
			v(r)
		}
	}

	return r
}

type OptionFunc func(*AliyunDrive)

func WithLogger(logger *logrus.Logger) OptionFunc {
	return func(ins *AliyunDrive) {
		ins.logger = logger
	}
}

func WithStore(store Store) OptionFunc {
	return func(ins *AliyunDrive) {
		token, err := store.Get(context.Background(), KeyAccessToken)
		if err == err && token != nil {
			ins.accessToken = string(token)
		}

		ins.store = store
	}
}
