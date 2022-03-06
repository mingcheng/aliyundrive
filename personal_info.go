package store

type PersonalInfoResp struct {
	PersonalRightsInfo struct {
		SpuId      string `json:"spu_id"`
		Name       string `json:"name"`
		IsExpires  bool   `json:"is_expires"`
		Privileges []struct {
			FeatureId     string `json:"feature_id"`
			FeatureAttrId string `json:"feature_attr_id"`
			Quota         int64  `json:"quota"`
		} `json:"privileges"`
	} `json:"personal_rights_info"`
	PersonalSpaceInfo struct {
		UsedSize  int   `json:"used_size"`
		TotalSize int64 `json:"total_size"`
	} `json:"personal_space_info"`
}
