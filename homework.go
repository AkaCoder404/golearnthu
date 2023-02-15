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
	SSearch              string     `json:"sSearch"`
	AaData               []Homework `json:"aaData"`
}
type Homework struct {
	Kssj    int    `json:"kssj"`    // 开始时间
	KssjStr string `json:"kssjStr"` // 开始时间 String
	Jzsj    int    `json:"jzsj"`    // 截止时间
	JzsjStr string `json:"jzsjStr"` // 截止时间 String
	Bt      string `json:"bt"`      // 标题
	Wz      int    `json:"wz"`      // 序号
	Zywcfs  int    `json:"zywcfs"`  // 作业完成方式
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

// UnsubmittedHomeworkInfo : A struct for a unsubmmited homework instance
type UnsubmittedHomeworkInfo struct {
	// 作业内容及要求
	Title                string // 作业标题
	Description          string // 作业说明
	AttachmentName       string // 作业附件
	AttachmentLink       string // 作业附件链接
	AttachmentSize       string // 作业附件大小
	AnswerDesc           string // 答案说明
	AnswerAttachmentName string // TODO 大答附件
	AnswerAttachmentLink string // TODO 答案附件链接
	AnswerAttachmentSize string // TODO 答案附件大小
	PublishObject        string // 发布对象
	FinishType           string // 完成方式
	DueDate              string // 截止日期
}

func (h *HomeworkService) GetUnsubmittedHomeworkInfo(courseID string, homeworkID string, xsHomeworkID string) (*UnsubmittedHomeworkInfo, error) {
	// Create request
	var emptyData url.Values = url.Values{}
	resp, err := h.client.Request(context.Background(), http.MethodPost, learnUnsubmittedHomeworkInfo(courseID, homeworkID, xsHomeworkID), strings.NewReader(emptyData.Encode()))
	if err != nil {
		return nil, err
	}

	// Handle response
	stringResponse := DecodeRequestBodyToString(resp)
	// stringResponse = strings.Split(stringResponse, "<div class=\"ttee\">作业内容及要求：</div>")[1]

	var response UnsubmittedHomeworkInfo

	// 作业标题
	response.Title = strings.Split(strings.Split(stringResponse, ">作业标题</div>")[1], "</div>")[0]
	response.Title = stripHtml(response.Title)

	// 作业说明
	response.Description = strings.Split(strings.Split(stringResponse, ">作业说明</div>")[1], "</div>")[0]
	response.Description = stripHtml(response.Description)

	// 作业附件
	response.AttachmentName = strings.Split(stringResponse, ">作业附件</div>")[1]
	response.AttachmentName = strings.Split(response.AttachmentName, "<!--答案说明-->")[0]
	// 处理未文件的情况
	if stripHtml(response.AttachmentName) != "" {
		attachment := response.AttachmentName
		response.AttachmentName = strings.Split(response.AttachmentName, "</a>")[0]
		response.AttachmentSize = strings.Split(attachment, "</span>")[1]
		response.AttachmentSize = stripHtml(response.AttachmentSize)
		response.AttachmentLink = strings.Split(attachment, "href=\"")[2]
		response.AttachmentLink = strings.Split(response.AttachmentLink, "\">")[0]
		response.AttachmentName = stripHtml(response.AttachmentName)
	} else {
		response.AttachmentName = ""
	}

	// 答案说明
	response.AnswerDesc = strings.Split(strings.Split(stringResponse, ">答案说明</div>")[1], "</div>")[0]
	response.AnswerDesc = stripHtml(response.AnswerDesc)

	// TODO 答案附件
	// response.AnswerAttachmentName = strings.Split(stringResponse, ">答案附件</div>")[1]

	// 发布对象
	response.PublishObject = strings.Split(strings.Split(stringResponse, ">发布对象</div>")[1], "</div>")[0]
	response.PublishObject = stripHtml(response.PublishObject)

	// 完成方式
	response.FinishType = strings.Split(strings.Split(stringResponse, ">完成方式</div>")[1], "</div>")[0]
	response.FinishType = stripHtml(response.FinishType)

	// 截止日期
	response.DueDate = strings.Split(strings.Split(stringResponse, ">截止日期(GMT+8)</div>")[1], "</div>")[0]
	response.DueDate = stripHtml(response.DueDate)
	return &response, nil
}

// SubmittedHomeworkInfo : A struct for a submitted homework, but not graded hw instance
type SubmittedHomeworkInfo struct {
	// 作业内容及要求
	Title          string // 作业标题
	Description    string // 作业说明
	AttachmentName string // 作业附件
	AttachmentLink string // 作业附件链接
	AttachmentSize string // TODO 作业附件大小

	AnswerDesc           string // 答案说明
	AnswerAttachmentName string // TODO 大答附件
	AnswerAttachmentLink string // TODO 答案附件链接
	AnswerAttachmentSize string // TODO 答案附件大小
	PublishObject        string // 发布对象
	FinishType           string // 完成方式

	// 本人提交的作业
	Studentid                string // 学号
	SubmissionDate           string // 提交日期
	SubmissionContent        string // 上交作业内容
	SubmissionAttachmentName string // 上交作业附件
	SubmissionAttachmentLink string // 上交作业附件连接
	SubmissionAttachmentSize string // 上交作业附件大小
}

// GetSubmittedHomeworkInfo : Get a submitted homework information
func (h *HomeworkService) GetSubmittedHomeworkInfo(courseID string, homeworkID string, xsHomeworkID string) (*SubmittedHomeworkInfo, error) {
	// Create request
	var emptyData url.Values = url.Values{}
	resp, err := h.client.Request(context.Background(), http.MethodPost, learnSubmittedNotGradedHWInfo(courseID, homeworkID, xsHomeworkID), strings.NewReader(emptyData.Encode()))
	if err != nil {
		return nil, err
	}

	// Handle response
	stringResponse := DecodeRequestBodyToString(resp)
	var response SubmittedHomeworkInfo

	// 作业标题
	response.Title = strings.Split(strings.Split(stringResponse, ">作业标题</div>")[1], "</div>")[0]
	response.Title = stripHtml(response.Title)

	// 作业说明
	response.Description = strings.Split(strings.Split(stringResponse, ">作业说明</div>")[1], "</div>")[0]
	response.Description = stripHtml(response.Description)

	// 作业附件
	response.AttachmentName = strings.Split(stringResponse, ">作业附件</div>")[1]
	response.AttachmentName = strings.Split(response.AttachmentName, "<!--答案说明-->")[0]
	// 处理未文件的情况
	if stripHtml(response.AttachmentName) != "" {
		attachment := response.AttachmentName
		response.AttachmentName = strings.Split(response.AttachmentName, "</a>")[0]
		response.AttachmentSize = strings.Split(attachment, "</span>")[1]
		response.AttachmentSize = stripHtml(response.AttachmentSize)
		response.AttachmentLink = strings.Split(attachment, "href=\"")[2]
		response.AttachmentLink = strings.Split(response.AttachmentLink, "\">")[0]
		response.AttachmentName = stripHtml(response.AttachmentName)
	} else {
		response.AttachmentName = ""
	}

	// 答案说明
	response.AnswerDesc = strings.Split(strings.Split(stringResponse, ">答案说明</div>")[1], "</div>")[0]
	response.AnswerDesc = stripHtml(response.AnswerDesc)

	// TODO 答案附件

	// 发布对象
	response.PublishObject = strings.Split(strings.Split(stringResponse, ">发布对象</div>")[1], "</div>")[0]
	response.PublishObject = stripHtml(response.PublishObject)

	// 完成方式
	response.FinishType = strings.Split(strings.Split(stringResponse, ">完成方式</div>")[1], "</div>")[0]
	response.FinishType = stripHtml(response.FinishType)

	// 学号
	response.Studentid = strings.Split(strings.Split(stringResponse, ">学号</div>")[1], "</div>")[0]
	response.Studentid = stripHtml(response.Studentid)

	// 提交日期
	response.SubmissionDate = strings.Split(strings.Split(stringResponse, ">提交日期</div>")[1], "</div>")[0]
	response.SubmissionDate = stripHtml(response.SubmissionDate)

	// 上交作业内容
	response.SubmissionContent = strings.Split(strings.Split(stringResponse, ">上交作业内容</div>")[1], "</div>")[0]
	response.SubmissionContent = stripHtml(response.SubmissionContent)

	// 上交作业附件
	response.SubmissionAttachmentName = strings.Split(stringResponse, ">上交作业附件</div>")[1]
	response.SubmissionAttachmentName = strings.Split(response.SubmissionAttachmentName, "下载")[0] // TODO handle no attachment
	// 处理未文件的情况
	if stripHtml(response.SubmissionAttachmentName) != "" {
		attachment := response.SubmissionAttachmentName
		response.SubmissionAttachmentName = strings.Split(response.SubmissionAttachmentName, "</a>")[0]
		response.SubmissionAttachmentSize = strings.Split(attachment, "</span>")[1]
		response.SubmissionAttachmentSize = stripHtml(response.SubmissionAttachmentSize)
		response.SubmissionAttachmentLink = strings.Split(attachment, "href=\"")[2]
		response.SubmissionAttachmentLink = strings.Split(response.SubmissionAttachmentLink, "\"")[0]
		response.SubmissionAttachmentName = stripHtml(response.SubmissionAttachmentName)
	} else {
		response.SubmissionAttachmentName = ""
	}

	return &response, nil
}

// GradedHomeworkInfo : A struct for a graded homework instance
type GradedHomeworkInfo struct {
	// 作业内容及要求
	Title                string // 作业标题
	Description          string // 作业说明
	AttachmentName       string // 作业附件
	AttachmentLink       string // 作业附件链接
	AttachmentSize       string // 作业附件大小
	AnswerDesc           string // 答案说明
	AnswerAttachment     string // 大答附件
	AnswerAttachmentLink string // TODO: 答案附件链接
	AnswerAttachmentSize string // TODO: 答案附件大小
	PublishObject        string // 发布对象
	FinishType           string // 完成方式
	DueDate              string // 截止日期

	// 本人提交的作业
	Studentid                string // 提交日期
	SubmissionDate           string // 提交日期
	SubmissionContent        string // 上交作业内容
	SubmissionAttachmentName string // 上交作业附件名字
	SubmissionAttachmentLink string // 上交作业附件连接
	SubmissionAttachmentSize string // TODO

	// 老师批阅结果
	ReviewTeacher          string // 批阅老师
	ReviewTime             string // 批阅时间
	Grade                  string // 成绩
	Comments               string // 评语
	CommentsAttachmentName string // 评语附件
	CommentsAttachmentLink string // TODO
	CommentsAttachmentSize string // TODO
}

func (h *HomeworkService) GetGradedHomeworkInfo(courseID string, homeworkID string, xsHomeworkID string) (*GradedHomeworkInfo, error) {
	var emptyData url.Values = url.Values{}
	resp, err := h.client.Request(context.Background(), http.MethodPost, learnGradedHomeworkInfo(courseID, homeworkID, xsHomeworkID), strings.NewReader(emptyData.Encode()))
	if err != nil {
		return nil, err
	}

	stringResponse := DecodeRequestBodyToString(resp)
	var response GradedHomeworkInfo

	// 作业标题
	response.Title = strings.Split(stringResponse, "<div class=\"left\">作业标题</div>")[1]
	response.Title = strings.Split(response.Title, "</p>")[0]
	response.Title = stripHtml(response.Title)

	// 作业说明
	response.Description = strings.Split(stringResponse, ">作业说明</div>")[1]
	response.Description = strings.Split(response.Description, "</div>")[0]
	response.Description = stripHtml(response.Description)

	// 作业附件
	response.AttachmentName = strings.Split(stringResponse, ">作业附件</div>")[1]
	response.AttachmentName = strings.Split(response.AttachmentName, "<!--答案说明-->")[0]
	// 处理未文件的情况
	if stripHtml(response.AttachmentName) != "" {
		attachment := response.AttachmentName
		response.AttachmentName = strings.Split(response.AttachmentName, "</a>")[0]
		response.AttachmentSize = strings.Split(attachment, "</span>")[1]
		response.AttachmentSize = stripHtml(response.AttachmentSize)
		response.AttachmentLink = strings.Split(attachment, "href=\"")[2]
		response.AttachmentLink = strings.Split(response.AttachmentLink, "\">")[0]
		response.AttachmentName = stripHtml(response.AttachmentName)
	} else {
		response.AttachmentName = ""
	}

	// Homework Answer Description
	response.AnswerDesc = strings.Split(stringResponse, ">答案说明</div>")[1]
	response.AnswerDesc = strings.Split(response.AnswerDesc, "</div>")[0]
	response.AnswerDesc = stripHtml(response.AnswerDesc)

	// Homework Answer Attachment
	// response.AnswerAttachment = strings.Split(stringResponse, ">答案附件</div>")[1]
	// response.AnswerAttachment = strings.Split(response.AnswerAttachment, "</div>")[0]
	// response.AnswerAttachment = stripHtml(response.AnswerAttachment)

	// Homework Publish Object
	response.PublishObject = strings.Split(stringResponse, "<div class=\"left\">发布对象</div>")[1]
	response.PublishObject = strings.Split(response.PublishObject, "</p>")[0]
	response.PublishObject = stripHtml(response.PublishObject)

	// Homework Finish Method
	response.FinishType = strings.Split(stringResponse, "<div class=\"left\">完成方式</div>")[1]
	response.FinishType = strings.Split(response.FinishType, "</p>")[0]
	response.FinishType = stripHtml(response.FinishType)

	// 作业截止时间
	response.DueDate = strings.Split(stringResponse, "<div class=\"left\">截止日期(GMT+8)</div>")[1]
	response.DueDate = strings.Split(response.DueDate, "</p>")[0]
	response.DueDate = stripHtml(response.DueDate)

	// 学生号
	response.Studentid = strings.Split(stringResponse, "<div class=\"left\">学号</div>")[1]
	response.Studentid = strings.Split(response.Studentid, "</p>")[0]
	response.Studentid = stripHtml(response.Studentid)

	// 作业上交日期
	response.SubmissionDate = strings.Split(stringResponse, "<div class=\"left\">提交日期</div>")[1]
	response.SubmissionDate = strings.Split(response.SubmissionDate, "</p>")[0]
	response.SubmissionDate = stripHtml(response.SubmissionDate)

	// 作业上交内容
	response.SubmissionContent = strings.Split(stringResponse, ">上交作业内容</div>")[1]
	response.SubmissionContent = strings.Split(response.SubmissionContent, "</div>")[0]
	response.SubmissionContent = stripHtml(response.SubmissionContent)

	// 上交作业附件
	response.SubmissionAttachmentName = strings.Split(stringResponse, ">上交作业附件</div>")[1]
	response.SubmissionAttachmentName = strings.Split(response.SubmissionAttachmentName, "老师批阅结果")[0]
	// 处理未文件的情况
	if stripHtml(response.SubmissionAttachmentName) != "" {
		attachment := response.SubmissionAttachmentName
		response.SubmissionAttachmentName = strings.Split(response.SubmissionAttachmentName, "</a>")[0]
		response.SubmissionAttachmentSize = strings.Split(attachment, "</span>")[1]
		response.SubmissionAttachmentSize = stripHtml(response.SubmissionAttachmentSize)
		response.SubmissionAttachmentLink = strings.Split(attachment, "href=\"")[2]
		response.SubmissionAttachmentLink = strings.Split(response.SubmissionAttachmentLink, "\"")[0]
		response.SubmissionAttachmentName = stripHtml(response.SubmissionAttachmentName)
	} else {
		response.SubmissionAttachmentName = ""
	}
	// Homework review teacher
	response.ReviewTeacher = strings.Split(stringResponse, "<div class=\"left\">批阅老师</div>")[1]
	response.ReviewTeacher = strings.Split(response.ReviewTeacher, "</p>")[0]
	response.ReviewTeacher = stripHtml(response.ReviewTeacher)

	// Homework review time
	response.ReviewTime = strings.Split(stringResponse, "<div class=\"left\">批阅时间</div>")[1]
	response.ReviewTime = strings.Split(response.ReviewTime, "</p>")[0]
	response.ReviewTime = stripHtml(response.ReviewTime)

	// Homework grade
	response.Grade = strings.Split(stringResponse, "<div class=\"left\">成绩</div>")[1]
	response.Grade = strings.Split(response.Grade, "</p>")[0]
	response.Grade = stripHtml(response.Grade)

	// Homework comments
	response.Comments = strings.Split(stringResponse, ">评语</div>")[1]
	response.Comments = strings.Split(response.Comments, "</span>")[0]
	response.Comments = stripHtml(response.Comments)

	// Homework comments attachment
	// response.CommentsAttachmentName = strings.Split(stringResponse, ">评语附件</div>")[1]
	// response.CommentsAttachmentName = strings.Split(response.CommentsAttachmentName, "</div>")[0]
	// response.CommentsAttachmentName = stripHtml(response.CommentsAttachmentName)

	return &response, nil
}
