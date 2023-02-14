package gothulearn

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type FileService service

// TODO: 课程文件 Tabs (FilePageList)
// FilePageList :
type FilePageList struct {
	Result string         `json:"result"`
	Msg    string         `json:"msg"`
	Object FilePageObject `json:"object"`
}

type FilePageObject struct {
	Page    int        `json:"page"`
	Total   int        `json:"total"`
	Records int        `json:"records"`
	Rows    []FilePage `json:"rows"`
}

type FilePage struct {
	Wlkcid string `json:"wlkcid"`
	Kjflid string `json:"kjflid"` //
	Lb     string `json:"lb"`
	Bt     string `json:"bt"` // 标题
	Wz     int    `json:"wz"` // 序号
	Xzst   int    `json:"xzst"`
	Czr    string `json:"czr"`
	Czsj   int    `json:"czsj"`
	Bz     string `json:"bz"`
	Id     string `json:"id"`
}

// TODO: 文件 per Tabs
type FileList struct {
	Result string `json:"result"`
	Msg    string `json:"msg"`
	Object []File `json:"object"`
}

// TODO: Better names for two, three, six, eight
type File struct {
	Id            string
	Filename      string // 文件名字
	Two           int
	Three         string
	Wlkcid        string // 文件ID
	Summary       string // 简要说明
	Six           string
	FileID        string // 文件ID
	Eight         int    //
	FileSize      int    // 文件大小Bytes
	PublishedDate string // 发布时间
}

// fileUnmarshaller :
func (f *File) UnmarshalJSON(data []byte) error {
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		fmt.Printf("Error while decoding %v\n", err)
		return err
	}
	// Unmarshall
	f.Id = v[0].(string)
	f.Filename = v[1].(string)
	f.Two = int(v[2].(float64))
	f.Three = v[3].(string)
	f.Wlkcid = v[4].(string)
	//
	if v[5] == nil {
		f.Summary = ""
	} else {
		f.Summary = v[5].(string)
	}
	f.Six = v[6].(string)
	f.FileID = v[7].(string)
	f.Eight = int(v[8].(float64))
	f.FileSize = int(v[9].(float64))
	f.PublishedDate = v[10].(string)
	return nil
}

// getFilePageList :
func (f *FileService) GetFilePageList(courseID string, classType string) (*FilePageList, error) {
	// Create request
	var emptyData url.Values = url.Values{}
	resp, err := f.client.Request(context.Background(), http.MethodPost, addCSRFTokenToUrl(learnFilePageList(courseID, "student"), f.client.csrf), strings.NewReader(emptyData.Encode()))
	if err != nil {
		return nil, err
	}

	// Handle response
	bytesResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response FilePageList
	json.Unmarshal([]byte(bytesResponse), &response)

	return &response, nil
}

// getFileList :
func (f *FileService) GetFileList(courseID string, filePageID string, classType string) (*FileList, error) {
	// Create request
	var emptyData url.Values = url.Values{}
	resp, err := f.client.Request(context.Background(), http.MethodPost, addCSRFTokenToUrl(learnFileList(courseID, filePageID, "student"), f.client.csrf), strings.NewReader(emptyData.Encode()))
	if err != nil {
		return nil, err
	}

	// Handle response
	bytesResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response FileList
	json.Unmarshal([]byte(bytesResponse), &response)

	return &response, nil
}

// downloadFile :
func (f *FileService) DownloadFileFromHomework(fileName string, courseID string, fileID string) (int64, error) {
	// Create request
	var emptyData url.Values = url.Values{}
	resp, err := f.client.Request(context.Background(), http.MethodGet, learnDownloadHWFile(courseID, fileID), strings.NewReader(emptyData.Encode()))
	if err != nil {
		return -1, err
	}

	// Create blank file
	filename := resp.Header["Content-Disposition"][0][len("attachment; filename="):]
	filename = strings.ReplaceAll(filename, "\"", "")
	file, err := os.Create(filename)
	if err != nil {
		return -1, err
	}

	// Put content on file and close
	size, err := io.Copy(file, resp.Body)
	defer file.Close()
	return size, nil
}

// downloadFileFromFilePage :
func (f *FileService) DownloadFileFromFilePage(fileName string, FileID string) (int64, error) {
	// Create request
	var emptyData url.Values = url.Values{}
	resp, err := f.client.Request(context.Background(), http.MethodGet, addCSRFTokenToUrl(learnDownloadFileFromPage(FileID), f.client.csrf), strings.NewReader(emptyData.Encode()))
	if err != nil {
		return -1, err
	}

	// Create blank file
	filename := resp.Header["Content-Disposition"][0][len("attachment; filename="):]
	filename = strings.ReplaceAll(filename, "\"", "")
	file, err := os.Create(filename)
	if err != nil {
		return -1, err
	}

	// Put content on file and close
	size, err := io.Copy(file, resp.Body)
	defer file.Close()
	return size, nil
}
