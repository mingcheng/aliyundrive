package aliyundrive

import (
	"context"
	"net/http"
)

type DeviceListResp struct {
	Result []struct {
		DeviceId   string `json:"deviceId"`
		DeviceName string `json:"deviceName"`
		ModelName  string `json:"modelName"`
		City       string `json:"city"`
		LoginTime  string `json:"loginTime"`
	} `json:"result"`
}

func (r *AliyunDrive) DeviceList(ctx context.Context) (*DeviceListResp, error) {
	var deviceListResp DeviceListResp

	_, err := r.request(ctx, &config{
		Method: http.MethodPost,
		URL:    "https://api.aliyundrive.com/users/v1/users/device_list",
		Body:   "{}",
	}, &deviceListResp)

	if err != nil {
		return nil, err
	}

	return &deviceListResp, nil
}
