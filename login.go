package aliyundrive

import "context"

// IsLogin to detect whether user is login
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
