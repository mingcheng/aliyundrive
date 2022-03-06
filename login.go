package aliyundrive

import "context"

func (r *AliyunDrive) IsLogin(ctx context.Context) bool {
	if r.accessToken == "" {
		return false
	}

	self, err := r.MySelf(ctx)
	if err != nil {
		return false
	}

	return self.UserID != ""
}
