package gothulearn

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type HomeworkService service

type HomeworkList struct {
	Result string         `json:"result"`
	Msg    string         `json:"msg"`
	Object HomeworkObject `json:"object"`
}

type HomeworkObject struct {
	ITotalRecords        string     `json:"iTotalRecords"`
	ITotalDisplayRecords string     `json:"iTotalDisplayRecords"`
	IDisplayStart        string     `json:"iDisplayStart"`
	IDisplayLength       string     `json:"iDisplayLength"`
	SEcho                string     `json:"sEcho"`
	SSearch              string     `json:"SSearch"`
	AaData               []Homework `json:"aaData"`
}

type Homework struct {
	Kssj    int    `json:"kssj"`    // 开始时间
	KssjStr string `json:"kssjStr"` // 开始时间 String
	Jzsj    int    `json:"jzsj"`    // 截止时间
	JzsjStr string `json:"jzsjStr"` // 截止时间 String
	Bt      string `json:"bt"`      // 标题
	Wz      int    `json:"wz"`      // 序号
	Zywcfs  int    `json:"zywcfs"`
	Zytjfs  int    `json:"zytjfs"`
	Jffs    int    `json:"jffs"`
	Mxdxmc  string `json:"mxdxmc"`
	Mxdx    string `json:"mxdx"`
	Wlkcid  string `json:"wlkcid"`
	Xszyid  string `json:"xszyid"`
	Xh      string `json:"xh"`      // 学号
	Zyid    string `json:"zyid"`    // 作业ID
	Zynr    string `json:"zynr"`    // 作业内容
	ZynrStr string `json:"zynrStr"` // 作业内容 Str
	Zyfjid  string `json:"zyfjid"`  // 作业附件ID
	Scr     string `json:"scr"`     // 上传人
	Scsj    string `json:"scsj"`    // 上传时间
	ScsjStr string `json:"scsjStr"` // 上传时间 Str
	Ggzh    string `json:"gzzh"`
	Pysj    string `json:"pysj"`    // 批阅时间
	PysjStr string `json:"pysjStr"` // 批阅时间 Str
	Pynr    string `json:"pynr"`    // 批阅内容
	Pylj    string `json:"pylj"`    //
	Cj      string `json:"cj"`      // 成绩
	Zt      string `json:"zt"`      // 状态
	Pyzt    string `json:"pyzt"`    // 批阅状态
	Qzid    string `json:"qzid"`
	Bz      string `json:"bz"`
	Qzmc    string `json:"qzmc"`
	Xm      string `json:"xm"` // 姓名
	Dwmc    string `json:"dwmc"`
	Bm      string `json:"bm"` // 班名
	Cjgbsj  string `json:"cjgbsj"`
	Fjdxxz  string `json:"fjdxxz"`
	Djzcj   string `json:"djzcj"`
	Ywdjzcj string `json:"ywdjzcj"`
	Jsm     string `json:"jsm"`  // 教师名
	Wjmc    string `json:"wjmc"` // 文件名称
	Wjdx    string `json:"wjdx"` // 文件大小
	Id      string `json:"id"`   // ID
}

func (h *HomeworkService) GetUnsubmittedHomeworks(courseID string) (*HomeworkList, error) {
	// Create request
	var payload url.Values = url.Values{}
	payload.Add("aoData", fmt.Sprintf("[{\"name\":\"sEcho\",\"value\":1},{\"name\":\"iColumns\",\"value\":8},{\"name\":\"sColumns\",\"value\":\",,,,,,,\"},{\"name\":\"iDisplayStart\",\"value\":0},{\"name\":\"iDisplayLength\",\"value\":\"30\"},{\"name\":\"mDataProp_0\",\"value\":\"wz\"},{\"name\":\"bSortable_0\",\"value\":false},{\"name\":\"mDataProp_1\",\"value\":\"bt\"},{\"name\":\"bSortable_1\",\"value\":true},{\"name\":\"mDataProp_2\",\"value\":\"mxdxmc\"},{\"name\":\"bSortable_2\",\"value\":true},{\"name\":\"mDataProp_3\",\"value\":\"zywcfs\"},{\"name\":\"bSortable_3\",\"value\":true},{\"name\":\"mDataProp_4\",\"value\":\"kssj\"},{\"name\":\"bSortable_4\",\"value\":true},{\"name\":\"mDataProp_5\",\"value\":\"jzsj\"},{\"name\":\"bSortable_5\",\"value\":true},{\"name\":\"mDataProp_6\",\"value\":\"jzsj\"},{\"name\":\"bSortable_6\",\"value\":true},{\"name\":\"mDataProp_7\",\"value\":\"function\"},{\"name\":\"bSortable_7\",\"value\":false},{\"name\":\"iSortCol_0\",\"value\":5},{\"name\":\"sSortDir_0\",\"value\":\"desc\"},{\"name\":\"iSortCol_1\",\"value\":6},{\"name\":\"sSortDir_1\",\"value\":\"desc\"},{\"name\":\"iSortingCols\",\"value\":2},{\"name\":\"wlkcid\",\"value\":\"%s\"}]", courseID))
	resp, err := h.client.Request(context.Background(), http.MethodPost, addCSRFTokenToUrl(learnUnsubmittedHomeworkList(courseID), h.client.csrf), strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, err
	}

	// Handle response
	bytesResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response HomeworkList
	json.Unmarshal(bytesResponse, &response)

	return &response, nil
}

func (h *HomeworkService) GetSubmmitedNotGradedHomeworks(courseID string) (*HomeworkList, error) {
	// Create request
	var payload url.Values = url.Values{}
	payload.Add("aoData", fmt.Sprintf("[{\"name\":\"sEcho\",\"value\":1},{\"name\":\"iColumns\",\"value\":8},{\"name\":\"sColumns\",\"value\":\",,,,,,,\"},{\"name\":\"iDisplayStart\",\"value\":0},{\"name\":\"iDisplayLength\",\"value\":\"30\"},{\"name\":\"mDataProp_0\",\"value\":\"wz\"},{\"name\":\"bSortable_0\",\"value\":false},{\"name\":\"mDataProp_1\",\"value\":\"bt\"},{\"name\":\"bSortable_1\",\"value\":true},{\"name\":\"mDataProp_2\",\"value\":\"mxdxmc\"},{\"name\":\"bSortable_2\",\"value\":true},{\"name\":\"mDataProp_3\",\"value\":\"kssj\"},{\"name\":\"bSortable_3\",\"value\":true},{\"name\":\"mDataProp_4\",\"value\":\"jzsj\"},{\"name\":\"bSortable_4\",\"value\":true},{\"name\":\"mDataProp_5\",\"value\":\"jzsj\"},{\"name\":\"bSortable_5\",\"value\":true},{\"name\":\"mDataProp_6\",\"value\":\"zyfjid\"},{\"name\":\"bSortable_6\",\"value\":false},{\"name\":\"mDataProp_7\",\"value\":\"function\"},{\"name\":\"bSortable_7\",\"value\":false},{\"name\":\"iSortCol_0\",\"value\":4},{\"name\":\"sSortDir_0\",\"value\":\"desc\"},{\"name\":\"iSortCol_1\",\"value\":5},{\"name\":\"sSortDir_1\",\"value\":\"desc\"},{\"name\":\"iSortingCols\",\"value\":2},{\"name\":\"wlkcid\",\"value\":\"%s\"}]", courseID))
	resp, err := h.client.Request(context.Background(), http.MethodPost, addCSRFTokenToUrl(learnSubmittedNotGradedHWList(courseID), h.client.csrf), strings.NewReader(payload.Encode()))
	// Handle response
	bytesResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response HomeworkList
	json.Unmarshal(bytesResponse, &response)

	return &response, nil
}

func (h *HomeworkService) GetGradedHomeworks(courseID string) (*HomeworkList, error) {
	// Create request
	var payload url.Values = url.Values{}
	payload.Add("aoData", fmt.Sprintf("[{\"name\":\"sEcho\",\"value\":1},{\"name\":\"iColumns\",\"value\":9},{\"name\":\"sColumns\",\"value\":\",,,,,,,,\"},{\"name\":\"iDisplayStart\",\"value\":0},{\"name\":\"iDisplayLength\",\"value\":\"30\"},{\"name\":\"mDataProp_0\",\"value\":\"wz\"},{\"name\":\"bSortable_0\",\"value\":true},{\"name\":\"mDataProp_1\",\"value\":\"bt\"},{\"name\":\"bSortable_1\",\"value\":true},{\"name\":\"mDataProp_2\",\"value\":\"zywcfs\"},{\"name\":\"bSortable_2\",\"value\":true},{\"name\":\"mDataProp_3\",\"value\":\"scsj\"},{\"name\":\"bSortable_3\",\"value\":true},{\"name\":\"mDataProp_4\",\"value\":\"jsm\"},{\"name\":\"bSortable_4\",\"value\":false},{\"name\":\"mDataProp_5\",\"value\":\"pysj\"},{\"name\":\"bSortable_5\",\"value\":true},{\"name\":\"mDataProp_6\",\"value\":\"cj\"},{\"name\":\"bSortable_6\",\"value\":true},{\"name\":\"mDataProp_7\",\"value\":\"zyfjid\"},{\"name\":\"bSortable_7\",\"value\":false},{\"name\":\"mDataProp_8\",\"value\":\"function\"},{\"name\":\"bSortable_8\",\"value\":false},{\"name\":\"iSortCol_0\",\"value\":0},{\"name\":\"sSortDir_0\",\"value\":\"desc\"},{\"name\":\"iSortingCols\",\"value\":1},{\"name\":\"wlkcid\",\"value\":\"%s\"}]", courseID))
	resp, err := h.client.Request(context.Background(), http.MethodPost, addCSRFTokenToUrl(learnGradedHomeworkList(courseID), h.client.csrf), strings.NewReader(payload.Encode()))
	// Handle response
	bytesResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response HomeworkList
	json.Unmarshal(bytesResponse, &response)

	return &response, nil
}
