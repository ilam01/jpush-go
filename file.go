package jpush

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
)

const (
	TypeAlias          = 1
	TypeRegistrationId = 2
)

type file struct {
	FileId     string `json:"file_id"`
	Type       string `json:"type"`
	CreateTime string `json:"create_time"`
}

type fileResponse struct {
	FileId string `json:"file_id"`
	ErrorResponse
}

type GetFilesResponse struct {
	TotalCount int    `json:"total_count"`
	Files      []file `json:"files"`
	ErrorResponse
}

// typeInt 1 别名 2 注册id 默认别名
func (j *JPush) SendFile(filePath string, typeInt int) (fileId string, err error) {
	//打开文件句柄操作
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	var typeString string
	switch typeInt {
	case TypeAlias:
		typeString = "alias"
	case TypeRegistrationId:
		typeString = "registration_id"
	default:
		typeString = "alias"
	}

	//创建一个模拟的form中的一个选项,这个form项现在是空的
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作, 设置文件的上传参数叫uploadfile, 文件名是filename,
	//相当于现在还没选择文件, form项里选择文件的选项
	fileWriter, err := bodyWriter.CreateFormFile("filename", "@"+typeString+".txt")
	if err != nil {
		return
	}

	//iocopy 这里相当于选择了文件,将文件放到form中
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return
	}

	//获取上传文件的类型,multipart/form-data; boundary=...
	contentType := bodyWriter.FormDataContentType()

	//这个很关键,必须这样写关闭,不能使用defer关闭,不然会导致错误
	bodyWriter.Close()

	//这里就是上传的其他参数设置,可以使用 bodyWriter.WriteField(key, val) 方法
	//也可以自己在重新使用  multipart.NewWriter 重新建立一项,这个再server 会有例子
	params := map[string]string{
		"filename": "@" + typeString + ".txt",
		"path":     filePath,
	}
	//这种设置值得仿佛 和下面再从新创建一个的一样
	for key, val := range params {
		_ = bodyWriter.WriteField(key, val)
	}

	//发送post请求到服务端
	respBody, err := j.requestFile("POST", j.GetURL("file")+typeString, contentType, bodyBuf)
	ret := new(fileResponse)
	err2 := json.Unmarshal(respBody, ret)
	if err2 != nil {
		return "", err
	}
	return ret.FileId, nil
}

func (j *JPush) GetFiles() (files *GetFilesResponse, err error) {
	url := j.GetURL("file")
	resp, err := j.request("GET", url, nil, nil)
	ret := new(GetFilesResponse)
	err2 := json.Unmarshal(resp, ret)
	if err2 != nil {
		return nil, err
	}
	return ret, err
}

func (j *JPush) DeleteFiles(fileId string) (delete bool, err error) {
	url := j.GetURL("file") + fileId
	_, err = j.request("DELETE", url, nil, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
