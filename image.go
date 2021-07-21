package jpush

import (
	"bytes"
	"encoding/json"
)

type IconSize int

const (
	BigPic IconSize = iota + 1
	LargeIcon
	SmallIcon
)

type ImagesByUrlsRequest struct {
	FcmImageURL     string   `json:"fcm_image_url,omitempty"`
	HuaweiImageURL  string   `json:"huawei_image_url,omitempty"`
	ImageType       IconSize `json:"image_type,omitempty"`
	ImageURL        string   `json:"image_url,omitempty"`
	JiguangImageURL string   `json:"jiguang_image_url,omitempty"`
	OppoImageURL    string   `json:"oppo_image_url,omitempty"`
	XiaomiImageURL  string   `json:"xiaomi_image_url,omitempty"`
}

type ImagesByUrlsResponse struct {
	FcmImageURL     string `json:"fcm_image_url,omitempty"`
	HuaweiImageURL  string `json:"huawei_image_url,omitempty"`
	JiguangImageURL string `json:"jiguang_image_url,omitempty"`
	MediaID         string `json:"media_id,omitempty"`
	OppoImageURL    string `json:"oppo_image_url,omitempty"`
	XiaomiImageURL  string `json:"xiaomi_image_url,omitempty"`
	ErrorResponse
}

// ImagesByUrls 通过图片地址获取媒体id
func (j *JPush) ImagesByUrls(req ImagesByUrlsRequest) (*ImagesByUrlsResponse, error) {
	url := j.GetURL("images") + "byurls"
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := j.request("POST", url, bytes.NewReader(buf), nil)
	ret := new(ImagesByUrlsResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err2
	}
	return ret, err
}

type ImagesByUrlsUpdateRequest struct {
	FcmImageURL     string `json:"fcm_image_url,omitempty"`
	HuaweiImageURL  string `json:"huawei_image_url,omitempty"`
	JiguangImageURL string `json:"jiguang_image_url,omitempty"`
	OppoImageURL    string `json:"oppo_image_url,omitempty"`
	XiaomiImageURL  string `json:"xiaomi_image_url,omitempty"`
}

// ImagesByFiles 通过图片文件获取媒体id
func (j *JPush) ImagesByFiles(media string, req ImagesByUrlsUpdateRequest) (*ImagesByUrlsResponse, error) {
	url := j.GetURL("images") + "/byurls/" + media
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := j.request("PUT", url, bytes.NewReader(buf), nil)
	ret := new(ImagesByUrlsResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err2
	}
	return ret, err
}
