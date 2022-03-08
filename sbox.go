package aliyundrive

import (
	"context"
	"fmt"
	"net/http"
)

type GetSBoxResp struct {
	DriveID          string `json:"drive_id"`
	UsedSize         int    `json:"sbox_used_size"`
	RealUsedSize     int    `json:"sbox_real_used_size"`
	TotalSize        int64  `json:"sbox_total_size"`
	RecommendVip     string `json:"recommend_vip"`
	PinSetup         bool   `json:"pin_setup"`
	Locked           bool   `json:"locked"`
	InsuranceEnabled bool   `json:"insurance_enabled"`
}

func (r *AliyunDrive) GetSBox(ctx context.Context) (*GetSBoxResp, error) {
	var result GetSBoxResp

	response, err := r.request(ctx, &config{
		Method: http.MethodPost,
		URL:    "https://api.aliyundrive.com/v2/sbox/get",
		Body:   "{}",
	}, &result)

	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, fmt.Errorf("%s", response.Status())
	}

	return &result, nil
}
