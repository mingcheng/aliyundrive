package aliyundrive

import (
	"context"
	"net/http"
)

type PersonalInfoResp struct {
	RightsInfo struct {
		SpuId      string `json:"spu_id"`
		Name       string `json:"name"`
		IsExpires  bool   `json:"is_expires"`
		Privileges []struct {
			FeatureId     string `json:"feature_id"`
			FeatureAttrId string `json:"feature_attr_id"`
			Quota         int    `json:"quota"`
		} `json:"privileges"`
	} `json:"personal_rights_info"`
	SpaceInfo struct {
		UsedSize  uint `json:"used_size"`
		TotalSize uint `json:"total_size"`
	} `json:"personal_space_info"`
}

func (r *AliyunDrive) PersonalInfo(ctx context.Context) (*PersonalInfoResp, error) {
	var personalInfoResp PersonalInfoResp

	_, err := r.request(ctx, &config{
		Method: http.MethodPost,
		URL:    "https://api.aliyundrive.com/v2/databox/get_personal_info",
		Body:   "{}",
	}, &personalInfoResp)

	if err != nil {
		return nil, err
	}

	return &personalInfoResp, nil
}
