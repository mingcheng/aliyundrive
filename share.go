package aliyundrive

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type ShareFile struct {
	Popularity    int       `json:"popularity"`
	ShareId       string    `json:"share_id"`
	ShareMsg      string    `json:"share_msg"`
	ShareName     string    `json:"share_name"`
	Description   string    `json:"description"`
	Expiration    string    `json:"expiration"`
	Expired       bool      `json:"expired"`
	SharePwd      string    `json:"share_pwd"`
	ShareUrl      string    `json:"share_url"`
	Creator       string    `json:"creator"`
	DriveId       string    `json:"drive_id"`
	FileId        string    `json:"file_id"`
	FileIdList    []string  `json:"file_id_list"`
	PreviewCount  int       `json:"preview_count"`
	SaveCount     int       `json:"save_count"`
	DownloadCount int       `json:"download_count"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	FirstFile     struct {
		Trashed            bool      `json:"trashed"`
		Category           string    `json:"category"`
		ContentHash        string    `json:"content_hash"`
		ContentHashName    string    `json:"content_hash_name"`
		ContentType        string    `json:"content_type"`
		Crc64Hash          string    `json:"crc64_hash"`
		CreatedAt          time.Time `json:"created_at"`
		DomainId           string    `json:"domain_id"`
		DownloadUrl        string    `json:"download_url"`
		DriveId            string    `json:"drive_id"`
		EncryptMode        string    `json:"encrypt_mode"`
		FileExtension      string    `json:"file_extension"`
		FileId             string    `json:"file_id"`
		Hidden             bool      `json:"hidden"`
		ImageMediaMetadata struct {
			CroppingSuggestion []struct {
				AspectRatio      string `json:"aspect_ratio"`
				CroppingBoundary struct {
					Height int `json:"height"`
					Left   int `json:"left"`
					Top    int `json:"top"`
					Width  int `json:"width"`
				} `json:"cropping_boundary"`
				Score float64 `json:"score"`
			} `json:"cropping_suggestion"`
			Exif         string `json:"exif"`
			Height       int    `json:"height"`
			ImageQuality struct {
				OverallScore float64 `json:"overall_score"`
			} `json:"image_quality"`
			ImageTags []struct {
				Confidence float64 `json:"confidence"`
				Name       string  `json:"name"`
				TagLevel   int     `json:"tag_level"`
				ParentName string  `json:"parent_name,omitempty"`
			} `json:"image_tags"`
			Width int `json:"width"`
		} `json:"image_media_metadata"`
		Labels       []string  `json:"labels"`
		MimeType     string    `json:"mime_type"`
		Name         string    `json:"name"`
		ParentFileId string    `json:"parent_file_id"`
		PunishFlag   int       `json:"punish_flag"`
		Size         int       `json:"size"`
		Starred      bool      `json:"starred"`
		Status       string    `json:"status"`
		Thumbnail    string    `json:"thumbnail"`
		Type         string    `json:"type"`
		UpdatedAt    time.Time `json:"updated_at"`
		UploadId     string    `json:"upload_id"`
		Url          string    `json:"url"`
		UserMeta     string    `json:"user_meta"`
	} `json:"first_file"`
	IsPhotoCollection bool   `json:"is_photo_collection"`
	SyncToHomepage    bool   `json:"sync_to_homepage"`
	PopularityStr     string `json:"popularity_str"`
	FullShareMsg      string `json:"full_share_msg"`
	DisplayName       string `json:"display_name"`
}

type CreateShareReq struct {
	Expiration     string   `json:"expiration"`
	SyncToHomepage bool     `json:"sync_to_homepage"`
	SharePwd       string   `json:"share_pwd"`
	DriveId        string   `json:"drive_id"`
	FileIdList     []string `json:"file_id_list"`
}

type CancelShareReq struct {
	ShareId string `json:"share_id"`
}

type UpdateShareReq struct {
	ShareId    string `json:"share_id"`
	Expiration string `json:"expiration"`
}

type ListShareReq struct {
	Creator         string `json:"creator"`
	IncludeCanceled bool   `json:"include_canceled"`
	Category        string `json:"category"`
	OrderBy         string `json:"order_by"`
	OrderDirection  string `json:"order_direction"`
}

type ListShareResp struct {
	Items      []ShareFile `json:"items"`
	NextMarker string      `json:"next_marker"`
}

func (r *AliyunDrive) CreateShare(ctx context.Context, req *CreateShareReq) (*ShareFile, error) {
	var result ShareFile

	response, err := r.request(ctx, &config{
		Method: http.MethodPost,
		URL:    "https://api.aliyundrive.com/adrive/v2/share_link/create",
		Body:   req,
	}, &result)

	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, fmt.Errorf("%s", response.Status())
	}

	return &result, nil
}

func (r *AliyunDrive) UpdateShare(ctx context.Context, req *UpdateShareReq) (*ShareFile, error) {
	var result ShareFile

	response, err := r.request(ctx, &config{
		Method: http.MethodPost,
		URL:    "https://api.aliyundrive.com/v2/share_link/update",
		Body:   req,
	}, &result)

	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, fmt.Errorf("%s", response.Status())
	}

	return &result, nil
}

func (r *AliyunDrive) ListShare(ctx context.Context, req *ListShareReq) (*ListShareResp, error) {
	var result ListShareResp

	response, err := r.request(ctx, &config{
		Method: http.MethodPost,
		URL:    "https://api.aliyundrive.com/adrive/v3/share_link/list",
		Body:   req,
	}, &result)

	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, fmt.Errorf("%s", response.Status())
	}

	return &result, nil
}

func (r *AliyunDrive) CancelShare(ctx context.Context, req *CancelShareReq) error {
	response, err := r.request(ctx, &config{
		Method: http.MethodPost,
		URL:    "https://api.aliyundrive.com/adrive/v2/share_link/cancel",
		Body:   req,
	}, nil)

	if err != nil {
		return err
	}

	if !response.IsSuccess() {
		return fmt.Errorf("%s", response.Status())
	}

	return nil
}
