package jpush

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// Platform define platform entry
type Platform struct {
	isAll     bool
	Platforms []string
}

// SetAll set isAll value
func (p *Platform) SetAll(all bool) {
	p.isAll = all
}

// UnmarshalJSON unmarshal json
func (p *Platform) UnmarshalJSON(data []byte) error {
	if string(data) == "all" {
		p.isAll = true
		return nil
	}
	p.isAll = false
	return json.Unmarshal(data, &p.Platforms)
}

// MarshalJSON marshal json
func (p *Platform) MarshalJSON() (data []byte, err error) {
	if p.isAll {
		return []byte(`"all"`), nil
	}
	return json.Marshal(p.Platforms)
}

// File define file
type File struct {
	FileId string `json:"file_id"`
}

// Audience define Audience
type Audience struct {
	Tag            []string `json:"tag,omitempty"`
	TagAnd         []string `json:"tag_and,omitempty"`
	TagNot         []string `json:"tag_not,omitempty"`
	Alias          []string `json:"alias,omitempty"`
	RegistrationID []string `json:"registration_id,omitempty"`
	Segment        []string `json:"segment,omitempty"`
	ABTest         []string `json:"abtest,omitempty"`
	File           *File    `json:"file,omitempty"`
}

// PushAudience define audience entry
type PushAudience struct {
	isAll bool
	Aud   *Audience
}

// SetAll set isAll
func (p *PushAudience) SetAll(all bool) {
	p.isAll = all
}

// UnmarshalJSON unmarshal json
func (p *PushAudience) UnmarshalJSON(data []byte) error {
	if string(data) == "all" {
		p.isAll = true
		return nil
	}
	p.isAll = false
	return json.Unmarshal(data, p.Aud)
}

// MarshalJSON marshal json
func (p *PushAudience) MarshalJSON() (data []byte, err error) {
	if p.isAll {
		return []byte(`"all"`), nil
	}
	return json.Marshal(p.Aud)
}

// PushNotification define notification
type PushNotification struct {
	Alert    string                `json:"alert,omitempty"`
	Android  *NotificationAndroid  `json:"android,omitempty"`
	IOS      *NotificationIOS      `json:"ios,omitempty"`
	WinPhone *NotificationWinPhone `json:"winphone,omitempty"`
}

// NotificationAndroid define android notification
type NotificationAndroid struct {
	Alert             string                 `json:"alert"`
	Title             string                 `json:"title,omitempty"`
	BuilderID         int                    `json:"builder_id,int,omitempty"`
	Priority          int                    `json:"priority,omitempty"`
	Category          string                 `json:"category,omitempty"`
	Style             int                    `json:"style,int,omitempty"`
	AlertType         int                    `json:"alert_type,int,omitempty"`
	BigText           string                 `json:"big_text,omitempty"`
	Inbox             map[string]interface{} `json:"inbox,omitempty"`
	BigPicPath        string                 `json:"big_pic_path,omitempty"`
	Extras            map[string]interface{} `json:"extras,omitempty"`
	LargeIcon         string                 `json:"large_icon,omitempty"`
	Intent            map[string]interface{} `json:"intent,omitempty"`
	UriAction         string                 `json:"uri_action,omitempty"`
	UriActivity       string                 `json:"uri_activity,omitempty"`
	BadgeAddNum       int                    `json:"badge_add_num,omitempty"`
	BadgeClass        string                 `json:"badge_class,omitempty"`
	Sound             string                 `json:"sound,omitempty"`
	ShowBeginTime     string                 `json:"show_begin_time,omitempty"`
	ShowEndTime       string                 `json:"show_end_time,omitempty"`
	DisplayForeground string                 `json:"display_foreground,omitempty"`
	SmallIconUri      string                 `json:"small_icon_uri,omitempty"`
}

// NotificationIOS define ios notification
type NotificationIOS struct {
	Alert            interface{}            `json:"alert"`
	Sound            string                 `json:"sound,omitempty"`
	Badge            int                    `json:"badge,int,omitempty"`
	ContentAvailable bool                   `json:"content-available,omitempty"`
	MutableContent   bool                   `json:"mutable-content,omitempty"`
	Category         string                 `json:"category,omitempty"`
	Extras           map[string]interface{} `json:"extras,omitempty"`
	ThreadId         string                 `json:"thread-id,omitempty"`
}

// NotificationWinPhone define winphone notification
type NotificationWinPhone struct {
	Alert    string                 `json:"alert"`
	Title    string                 `json:"title,omitempty"`
	OpenPage string                 `json:"_open_page,omitempty"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

// PushMessage define push message
type PushMessage struct {
	MsgContent  string                 `json:"msg_content"`
	Title       string                 `json:"title,omitempty"`
	ContentType string                 `json:"content_type,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

// SmsMessage define sms message
type SmsMessage struct {
	DelayTime int                    `json:"delay_time,int"`
	TempID    float64                `json:"temp_id,float"`
	TempPara  map[string]interface{} `json:"temp_para,omitempty"`
}

// PushOptions define options
type PushOptions struct {
	SendNo            int               `json:"sendno,int,omitempty"`
	TimeToLive        int               `json:"time_to_live,int,omitempty"`
	OverrideMsgID     int64             `json:"override_msg_id,int64,omitempty"`
	ApnsProduction    bool              `json:"apns_production"`
	ApnsCollapseID    string            `json:"apns_collapse_id,omitempty"`
	BigPushDuration   int               `json:"big_push_duration,int,omitempty"`
	ThirdPartyChannel ThirdPartyChannel `json:"third_party_channel,omitempty"`
}

type ThirdPartyChannel struct {
	Fcm    Fcm    `json:"fcm,omitempty"`
	Huawei Huawei `json:"huawei,omitempty"`
	Meizu  Meizu  `json:"meizu,omitempty"`
	Oppo   Oppo   `json:"oppo,omitempty"`
	Vivo   Vivo   `json:"vivo,omitempty"`
	Xiaomi Xiaomi `json:"xiaomi,omitempty"`
}

type Fcm struct {
	Distribution    string `json:"distribution,omitempty"`
	DistributionFcm string `json:"distribution_fcm,omitempty"`
}
type Huawei struct {
	Fcm
	Importance         string `json:"importance,omitempty"`
	Inbox              string `json:"inbox,omitempty"`
	LargeIcon          string `json:"large_icon,omitempty"`
	OnlyUseVendorStyle bool   `json:"only_use_vendor_style,omitempty"`
	SmallIconURI       string `json:"small_icon_uri,omitempty"`
	Style              int64  `json:"style,omitempty"`
}
type Meizu struct {
	Fcm
}
type Oppo struct {
	Fcm
	BigPicPath string `json:"big_pic_path,omitempty"`
	ChannelID  string `json:"channel_id,omitempty"`
	LargeIcon  string `json:"large_icon,omitempty"`
	Style      int64  `json:"style,omitempty"`
}
type Vivo struct {
	Fcm
	Classification int64 `json:"classification,omitempty"`
	PushMode       int64 `json:"push_mode,omitempty"`
}
type Xiaomi struct {
	Fcm
	BigText               string `json:"big_text,omitempty"`
	ChannelID             string `json:"channel_id,omitempty"`
	DistributionCustomize string `json:"distribution_customize,omitempty"`
	LargeIcon             string `json:"large_icon,omitempty"`
	SmallIconColor        string `json:"small_icon_color,omitempty"`
	SmallIconURI          string `json:"small_icon_uri,omitempty"`
	Style                 int64  `json:"style,omitempty"`
}

const ToDevice int32 = 1
const ToUser int32 = 2

type CallbackParam struct {
	To        string `json:"to,omitempty"`
	ToType    int32  `json:"to_type,omitempty"`
	MessageId string `json:"message_id,omitempty"`
	Href      string `json:"href,omitempty"`
	Crowd     string `json:"crowd"`
}

type Callback struct {
	Url    string        `json:"url,omitempty"`
	Params CallbackParam `json:"params,omitempty"`
	Type   string        `json:"type,omitempty"`
}

// PushRequest define push request body
type PushRequest struct {
	Cid          string            `json:"cid,omitempty"`
	Platform     *Platform         `json:"platform"`
	Audience     *PushAudience     `json:"audience"`
	Notification *PushNotification `json:"notification,omitempty"`
	Message      *PushMessage      `json:"message,omitempty"`
	SmsMessage   *SmsMessage       `json:"sms_message,omitempty"`
	Options      *PushOptions      `json:"options,omitempty"`
	Callback     *Callback         `json:"callback,omitempty"`
}

// PushResponse define push repsone
type PushResponse struct {
	ErrorResponse
	MsgID  string `json:"msg_id"`
	Sendno string `json:"sendno"`
}

// PushCIDResponse get cid response
type PushCIDResponse struct {
	ErrorResponse
	Cids []string `json:"cidlist"`
}

// Push push notification or message to devices
// POST /v3/push
func (j *JPush) Push(req *PushRequest) (*PushResponse, error) {
	url := j.GetURL("push") + "push"
	if req.Audience.Aud.File != nil {
		url += "/file"
	}
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := j.request("POST", url, bytes.NewReader(buf), nil)
	ret := new(PushResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err2
	}
	return ret, err
}

// ScheduleGetCid get push by cid
// GET /v3/push/cid[?count=n[&type=xx]]
func (j *JPush) ScheduleGetCid() (*PushCIDResponse, error) {
	return j.PushGetCids(1, "schedule")
}

// PushGetCid get push by cid
// GET /v3/push/cid[?count=n[&type=xx]]
func (j *JPush) PushGetCid() (*PushCIDResponse, error) {
	return j.PushGetCids(1, "push")
}

// PushGetCids get push by cid
// GET /v3/push/cid[?count=n[&type=xx]]
func (j *JPush) PushGetCids(count int, cidtype string) (*PushCIDResponse, error) {
	url := j.GetURL("push") + "push/cid"
	params := make(map[string]string)
	params["count"] = strconv.Itoa(count)
	params["type"] = cidtype

	resp, err := j.request("GET", url, nil, params)
	ret := new(PushCIDResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err
	}
	return ret, err
}

// GroupPush group push
// POST /v3/grouppush
func (j *GroupPush) GroupPush(req *PushRequest) (map[string]PushResponse, error) {
	url := j.GetURL("push") + "grouppush"
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := j.request("POST", url, bytes.NewReader(buf), nil)
	if err != nil {
		return nil, err
	}
	ret := new(map[string]PushResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err
	}
	return *ret, err
}

// PushValidate push validate, not real push
// POST /v3/push/validate
func (j *JPush) PushValidate(req *PushRequest) (*PushResponse, error) {
	url := j.GetURL("push") + "push/validate"
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := j.request("POST", url, bytes.NewReader(buf), nil)
	ret := new(PushResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err
	}
	return ret, err
}
