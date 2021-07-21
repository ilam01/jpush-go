package jpush

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
)

// ReportTime report time
type ReportTime time.Time

// UnmarshalJSON unmarshal json
func (r *ReportTime) UnmarshalJSON(data []byte) error {
	t, err := time.Parse("2006-01-02", string(data))
	if err != nil {
		return err
	}
	*r = ReportTime(t)
	return nil
}

// MarshalJSON marshal json
func (r *ReportTime) MarshalJSON() (data []byte, err error) {
	t := time.Time(*r).Format("2006-01-02")
	return []byte(t), nil
}

// ReportStatusRequest push message status
type ReportStatusRequest struct {
	MsgID           int         `json:"msg_id,int"`
	RegistrationIds []string    `json:"registration_ids"`
	Date            *ReportTime `json:"date,omitempty"`
}

// ReportReceivedResponse define report received
type ReportReceivedResponse struct {
	ErrorResponse
	MsgID           float64 `json:"msg_id"`
	AndroidReceived int     `json:"android_received"`
	IOSApnsSent     int     `json:"ios_apns_sent"`
	IOSApnsReceived int     `json:"ios_apns_received"`
	IOSMsgReceived  int     `json:"ios_msg_received"`
	WpMpnsSent      int     `json:"wp_mpns_sent"`
}

// MessageStatus message status
type MessageStatus struct {
	Status int `json:"status"`
}

// ReportAndroidMessage android message
type ReportAndroidMessage struct {
	Received   int `json:"received,omitempty"`
	Target     int `json:"target,omitempty"`
	OnlinePush int `json:"online_push,omitempty"`
	Click      int `json:"click,omitempty"`
	MsgClick   int `json:"msg_click,omitempty"`
}

// ReportIOSMessage ios message
type ReportIOSMessage struct {
	ApnsSent     int `json:"apns_sent,omitempty"`
	ApnsTarget   int `json:"apns_target,omitempty"`
	ApnsReceived int `json:"apns_received,omitempty"`
	Click        int `json:"click,omitempty"`
	Target       int `json:"target,omitempty"`
	Received     int `json:"received,omitempty"`
}

// ReportWpMessage wp message
type ReportWpMessage struct {
	MpnsTarget int `json:"mpns_target,omitempty"`
	MpnsSent   int `json:"mpns_sent,omitempty"`
	Click      int `json:"click,omitempty"`
}

// ReportMessagesResponse define report messages
type ReportMessagesResponse struct {
	ErrorResponse
	MsgID   string                `json:"msg_id,omitempty"`
	Android *ReportAndroidMessage `json:"android,omitempty"`
	IOS     *ReportIOSMessage     `json:"ios,omitempty"`
	Wp      *ReportWpMessage      `json:"winphone,omitempty"`
}

// ReportUserAndroid define report android user
type ReportUserAndroid struct {
	New    int `json:"new,omitempty"`
	Active int `json:"active,omitempty"`
	Online int `json:"online,omitempty"`
}

// ReportUserIOS define report ios user
type ReportUserIOS struct {
	New    int `json:"new,omitempty"`
	Active int `json:"active,omitempty"`
	Online int `json:"online,omitempty"`
}

// ReportUser define report user item
type ReportUser struct {
	Time    *ReportTime        `json:"time"`
	Android *ReportUserAndroid `json:"android,omitempty"`
	IOS     *ReportUserIOS     `json:"ios,omitempty"`
}

// ReportDetailResponse message statistics details
type ReportDetailResponse struct {
	MsgId string `json:"msg_id"`
	Jpush struct {
		Target     int `json:"target"`
		OnlinePush int `json:"online_push"`
		Received   int `json:"received"`
		Click      int `json:"click"`
		MsgClick   int `json:"msg_click"`
	} `json:"jpush"`
	AndroidPns struct {
		PnsTarget   int `json:"pns_target"`
		PnsSent     int `json:"pns_sent"`
		PnsReceived int `json:"pns_received"`
		XmDetail    struct {
			Target   int `json:"target"`
			Sent     int `json:"sent"`
			Received int `json:"received"`
		} `json:"xm_detail"`
		HwDetail struct {
			Target   int `json:"target"`
			Sent     int `json:"sent"`
			Received int `json:"received"`
		} `json:"hw_detail"`
		MzDetail struct {
			Target   int `json:"target"`
			Sent     int `json:"sent"`
			Received int `json:"received"`
		} `json:"mz_detail"`
		OppoDetail struct {
			Target   int `json:"target"`
			Sent     int `json:"sent"`
			Received int `json:"received"`
		} `json:"oppo_detail"`
		VivoDetail struct {
			Target   int `json:"target"`
			Sent     int `json:"sent"`
			Received int `json:"received"`
		} `json:"vivo_detail"`
		FcmDetail struct {
			Target   int `json:"target"`
			Sent     int `json:"sent"`
			Received int `json:"received"`
		} `json:"fcm_detail"`
	} `json:"android_pns"`
	Ios struct {
		ApnsTarget   int `json:"apns_target"`
		ApnsSent     int `json:"apns_sent"`
		ApnsReceived int `json:"apns_received"`
		ApnsClick    int `json:"apns_click"`
		MsgTarget    int `json:"msg_target"`
		MsgReceived  int `json:"msg_received"`
		MsgClick     int `json:"msg_click"`
	} `json:"ios"`
	Winphone struct {
		MpnsTarget int `json:"mpns_target"`
		MpnsSent   int `json:"mpns_sent"`
		Click      int `json:"click"`
	} `json:"winphone"`
	QuickappJpush struct {
		Target     int `json:"target"`
		OnlinePush int `json:"online_push"`
		Received   int `json:"received"`
		Click      int `json:"click"`
		MsgClick   int `json:"msg_click"`
	} `json:"quickapp_jpush"`
	QuickappPns struct {
		PnsTarget int `json:"pns_target"`
		PnsSent   int `json:"pns_sent"`
	} `json:"quickapp_pns"`
}

// ReportUsersResponse define report user response
type ReportUsersResponse struct {
	ErrorResponse
	TimeUnit string       `json:"time_unit"`
	Start    *ReportTime  `json:"start"`
	Duration int          `json:"duration"`
	Items    []ReportUser `json:"items"`
}

// ReportReceived report received
// GET /v3/received
// Deprecated: it is old func
func (j *JPush) ReportReceived(msgIds []string) ([]ReportReceivedResponse, error) {
	url := j.GetURL("report") + "received"
	params := make(map[string]string)
	params["msg_ids"] = strings.Join(msgIds, ",")

	resp, err := j.request("GET", url, nil, params)
	var ret []ReportReceivedResponse
	err2 := json.Unmarshal(resp, &ret)
	if err2 != nil {
		return nil, err
	}
	return ret, err
}

// ReportStatus report push status
// POST /v3/status/message
func (j *JPush) ReportStatus(req *ReportStatusRequest) (map[string]MessageStatus, error) {
	url := j.GetURL("report") + "status/message"
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := j.request("POST", url, bytes.NewReader(buf), nil)
	if err != nil {
		return nil, err
	}
	ret := new(map[string]MessageStatus)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err
	}
	return *ret, err
}

// ReportMessages message stat
// GET /v3/messages
// Deprecated: it is old func
func (j *JPush) ReportMessages(msgIds []string) (*ReportMessagesResponse, error) {
	url := j.GetURL("report") + "messages"
	params := make(map[string]string)
	params["msg_ids"] = strings.Join(msgIds, ",")

	resp, err := j.request("GET", url, nil, params)
	ret := new(ReportMessagesResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err
	}
	return ret, err
}

// ReportUsers user stat
// GET /v3/users
func (j *JPush) ReportUsers(timeUnit string, start time.Time, duration int) (*ReportUsersResponse, error) {
	url := j.GetURL("report") + "users"
	params := make(map[string]string)
	params["time_unit"] = timeUnit
	params["start"] = ""
	if timeUnit == "HOUR" {
		params["start"] = start.Format("2006-01-02 15")
	} else if timeUnit == "DAY" {
		params["start"] = start.Format("2006-01-02")
	} else if timeUnit == "MONTH" {
		params["start"] = start.Format("2006-01")
	} else {
		return nil, errors.New("Bad Request: wrong time unit")
	}
	params["duration"] = strconv.Itoa(duration)

	resp, err := j.request("GET", url, nil, params)
	ret := new(ReportUsersResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err
	}
	return ret, err
}

// ReportMessagesDetail report messages detail
// GET /v3/messages/detail
func (j *JPush) ReportMessagesDetail(msgIds []string) ([]ReportDetailResponse, error) {
	url := j.GetURL("report") + "messages/detail"
	params := make(map[string]string)
	params["msg_ids"] = strings.Join(msgIds, ",")

	resp, err := j.request("GET", url, nil, params)
	var ret []ReportDetailResponse
	err2 := json.Unmarshal(resp, &ret)
	if err2 != nil {
		return nil, err
	}
	return ret, err
}
