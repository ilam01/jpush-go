package jpush

import (
	"bytes"
	"encoding/json"
	"strings"
)

// DeviceModify add and remove entry
type DeviceModify struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
}

// DeviceRegistrationIDRequest modify device request
type DeviceRegistrationIDRequest struct {
	Tags   *DeviceModify `json:"tags,omitempty"`
	Alias  string        `json:"alias,omitempty"`
	Mobile string        `json:"mobile,omitempty"`
}

// DeviceTagsRequest modify device named tag request
type DeviceTagsRequest struct {
	RegistrationIDs *DeviceModify `json:"registration_ids,omitempty"`
}

// DeviceStatusRequest get user status
type DeviceStatusRequest struct {
	RegistrationIDs []string `json:"registration_ids,omitempty"`
}

// DeviceRegistrationIDResponse get registration id response
type DeviceRegistrationIDResponse struct {
	ErrorResponse
	Tags   []string `json:"tags"`
	Alias  string   `json:"alias"`
	Mobile string   `json:"mobile,omitempty"`
}

// DeviceAliasResponse get alias response
type DeviceAliasResponse struct {
	ErrorResponse
	RegistrationIDs []string `json:"registration_ids"`
}

// DeviceTagsListResponse get tags list response
type DeviceTagsListResponse struct {
	ErrorResponse
	Tags []string `json:"tags"`
}

// DeviceTagsRegistrationIDResponse tags bind registration response
type DeviceTagsRegistrationIDResponse struct {
	ErrorResponse
	Result bool `json:"result"`
}

// DeviceStatusResponse get device status response
type DeviceStatusResponse struct {
	ErrorResponse
	Online         bool   `json:"online"`
	LastOnlineTime string `json:"last_online_time,omitempty"`
}

// DeviceGetRegistrationID get device info
func (j *JPush) DeviceGetRegistrationID(registrationID string) (*DeviceRegistrationIDResponse, error) {
	url := j.GetURL("device") + registrationID

	resp, err := j.request("GET", url, nil, nil)
	ret := new(DeviceRegistrationIDResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err
	}
	return ret, err
}

// DevicePostRegistrationID modify device info
func (j *JPush) DevicePostRegistrationID(registrationID string, req *DeviceRegistrationIDRequest) (*DefaultResponse, error) {
	url := j.GetURL("device") + registrationID
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := j.request("POST", url, bytes.NewReader(buf), nil)
	ret := new(DefaultResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err
	}
	return ret, err
}

// DeviceGetAlias get device named alias
func (j *JPush) DeviceGetAlias(alias string, platforms []string) (*DeviceAliasResponse, error) {
	url := j.GetURL("alias") + alias
	params := make(map[string]string)
	if len(platforms) > 0 {
		params["platform"] = strings.Join(platforms, ",")
	}

	resp, err := j.request("GET", url, nil, params)
	ret := new(DeviceAliasResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err
	}
	return ret, err
}

// DeviceDeleteAlias delete alias
func (j *JPush) DeviceDeleteAlias(alias string, platforms []string) (*DefaultResponse, error) {
	url := j.GetURL("alias") + alias
	params := make(map[string]string)
	if len(platforms) > 0 {
		params["platform"] = strings.Join(platforms, ",")
	}

	resp, err := j.request("DELETE", url, nil, params)
	ret := new(DefaultResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err
	}
	return ret, err
}

// DeviceGetTags get all tag list
func (j *JPush) DeviceGetTags() (*DeviceTagsListResponse, error) {
	url := j.GetURL("tag")

	resp, err := j.request("GET", url, nil, nil)
	ret := new(DeviceTagsListResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err
	}
	return ret, err
}

// DeviceGetTagsRegistrationID get device tag
func (j *JPush) DeviceGetTagsRegistrationID(tag string, registrationID string) (*DeviceTagsRegistrationIDResponse, error) {
	url := j.GetURL("tag") + tag + "/registration_ids/" + registrationID

	resp, err := j.request("GET", url, nil, nil)
	ret := new(DeviceTagsRegistrationIDResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err
	}
	return ret, err
}

// DevicePostTags modify device tag
func (j *JPush) DevicePostTags(tag string, req *DeviceTagsRequest) (*DefaultResponse, error) {
	url := j.GetURL("tag") + tag
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := j.request("POST", url, bytes.NewReader(buf), nil)
	ret := new(DefaultResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err
	}
	return ret, err
}

// DeviceDeleteTags delete tag
func (j *JPush) DeviceDeleteTags(tag string, platforms []string) (*DefaultResponse, error) {
	url := j.GetURL("tag") + tag
	params := make(map[string]string)
	if len(platforms) > 0 {
		params["platform"] = strings.Join(platforms, ",")
	}

	resp, err := j.request("DELETE", url, nil, params)
	ret := new(DefaultResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err
	}
	return ret, err
}

// DevicePostStatus get devices status
func (j *JPush) DevicePostStatus(req *DeviceStatusRequest) (map[string]DeviceStatusResponse, error) {
	url := j.GetURL("device") + "status/"
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := j.request("POST", url, bytes.NewReader(buf), nil)
	if err != nil {
		return nil, err
	}
	ret := new(map[string]DeviceStatusResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err
	}
	return *ret, err
}
