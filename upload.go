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
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

const maxPartSize = 1024 * 1024 * 1024 // 每个分片的大小

type UploadFileReq struct {
	DriveID       string `json:"drive_id"`
	ParentID      string `json:"parent_id"`
	FilePath      string `json:"file_path"`
	CheckNameMode string `json:"check_name_mode"`
	Name          string `json:"name"`
}

type UploadFileResp struct {
	DriveID            string    `json:"drive_id"`
	DomainID           string    `json:"domain_id"`
	FileID             string    `json:"file_id"`
	Name               string    `json:"name"`
	Type               string    `json:"type"`
	ContentType        string    `json:"content_type"` // application/oct-stream
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	FileExtension      string    `json:"file_extension"`
	Hidden             bool      `json:"hidden"`
	Size               int       `json:"size"`
	Starred            bool      `json:"starred"`
	Status             string    `json:"status"` // available
	UploadID           string    `json:"upload_id"`
	ParentFileID       string    `json:"parent_file_id"`
	Crc64Hash          string    `json:"crc64_hash"`
	ContentHash        string    `json:"content_hash"`
	ContentHashName    string    `json:"content_hash_name"` // sha1
	Category           string    `json:"category"`
	EncryptMode        string    `json:"encrypt_mode"`
	ImageMediaMetadata struct {
		ImageQuality struct{} `json:"image_quality"`
	} `json:"image_media_metadata"`
	Location string `json:"location"`
}

type completeUploadReq struct {
	DriveID  string `json:"drive_id"`
	UploadID string `json:"upload_id"`
	FileID   string `json:"file_id"`
}

type createFileWithProofReq struct {
	DriveID       string      `json:"drive_id"`
	PartInfoList  []*partInfo `json:"part_info_list"`
	ParentFileID  string      `json:"parent_file_id"`
	Name          string      `json:"name"`
	Type          string      `json:"type"`
	CheckNameMode string      `json:"check_name_mode"`
	Size          int64       `json:"size"`
	PreHash       string      `json:"pre_hash"`
}

type partInfo struct {
	PartNumber        int    `json:"part_number"`
	UploadURL         string `json:"upload_url"`
	InternalUploadURL string `json:"internal_upload_url"`
	ContentType       string `json:"content_type"`
}

type createFileWithProofResp struct {
	Type         string      `json:"type"`
	RapidUpload  bool        `json:"rapid_upload"`
	DomainId     string      `json:"domain_id"`
	DriveId      string      `json:"drive_id"`
	FileName     string      `json:"file_name"`
	EncryptMode  string      `json:"encrypt_mode"`
	Location     string      `json:"location"`
	UploadID     string      `json:"upload_id"`
	FileID       string      `json:"file_id"`
	PartInfoList []*partInfo `json:"part_info_list,omitempty"`
}

func makePartInfoList(size int64) []*partInfo {
	var res []*partInfo

	partInfoNum := int(size / maxPartSize)
	if size%maxPartSize > 0 {
		partInfoNum += 1
	}

	for i := 0; i < partInfoNum; i++ {
		res = append(res, &partInfo{PartNumber: i + 1})
	}

	return res
}

func (r *AliyunDrive) UploadFile(ctx context.Context, request *UploadFileReq) (*UploadFileResp, error) {
	file, err := os.Open(request.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if fileInfo.IsDir() {
		// TODO：支持文件夹
		return nil, fmt.Errorf("unsupport dir upload")
	}

	if request.CheckNameMode == "" {
		request.CheckNameMode = ModeRefuse
	}

	return r.UploadStream(ctx, request, file, fileInfo.Size())
}

func (r *AliyunDrive) UploadStream(ctx context.Context, request *UploadFileReq, stream io.Reader, streamSize int64) (*UploadFileResp, error) {
	proofResp, err := r.createFileWithProof(ctx, &createFileWithProofReq{
		DriveID:       request.DriveID,
		PartInfoList:  makePartInfoList(streamSize),
		ParentFileID:  request.ParentID,
		Name:          path.Base(request.FilePath),
		Type:          TypeFile,
		CheckNameMode: request.CheckNameMode,
		Size:          streamSize,
		PreHash:       "",
	})

	if err != nil {
		return nil, err
	}

	for _, part := range proofResp.PartInfoList {
		// TODO: 并发？
		err = r.uploadPart(ctx, part.UploadURL, io.LimitReader(stream, maxPartSize))
		if err != nil {
			return nil, err
		}
	}

	return r.completeUpload(ctx, &completeUploadReq{
		DriveID:  request.DriveID,
		UploadID: proofResp.UploadID,
		FileID:   proofResp.FileID,
	})
}

func (r *AliyunDrive) createFileWithProof(ctx context.Context, request *createFileWithProofReq) (*createFileWithProofResp, error) {
	var result createFileWithProofResp

	response, err := r.request(ctx, &config{
		Method: http.MethodPost,
		URL:    "https://api.aliyundrive.com/adrive/v2/file/create_with_proof",
		Body:   request,
	}, &result)
	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, fmt.Errorf("%s", response.Status())
	}

	return &result, nil
}

func (r *AliyunDrive) uploadPart(_ context.Context, uri string, reader io.Reader) error {
	req, err := http.NewRequest(http.MethodPut, uri, reader)

	response, err := r.client.GetClient().Do(req)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", response.Status)
	}

	return nil
}

func (r *AliyunDrive) completeUpload(ctx context.Context, request *completeUploadReq) (*UploadFileResp, error) {
	var result UploadFileResp

	_, err := r.request(ctx, &config{
		Method: http.MethodPost,
		URL:    "https://api.aliyundrive.com/adrive/v2/file/complete",
		Body:   request,
	}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
