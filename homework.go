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

// UnsubmittedHomeworkInfo : A struct for a unsubmmited homework instance
type UnsubmittedHomeworkInfo struct {
	Title            string // 作业标题
	Description      string // 作业说明
	Attachment       string // 作业附件
	AnswerDesc       string // 答案说明
	AnswerAttachment string // 大答附件
	PublishObject    string // 发布对象
	FinishType       string // 完成方式
	DueDate          string // 截止日期
}

// SubmittedHomeworkInfo : A struct for a submitted homework, but not graded hw instance
type SubmittedHomeworkInfo struct {
	// 作业内容及要求
	Title            string // 作业标题
	Description      string // 作业说明
	Attachment       string // 作业附件
	AnswerDesc       string // 答案说明
	AnswerAttachment string // 大答附件
	PublishObject    string // 发布对象
	FinishType       string // 完成方式
	DueDate          string // 截止日期

	// 本人提交的作业
	Studentid            string // 学号
	SubmissionDate       string // 提交日期
	SubmissionContent    string // 上交作业内容
	SubmissionAttachment string // 上交作业附件
}

// GradedHomeworkInfo :
type GradedHomeworkInfo struct {
	// 作业内容及要求
	Title                string // 作业标题
	Description          string // 作业说明
	Attachment           string // 作业附件
	AnswerDesc           string // 答案说明
	AnswerAttachment     string // 大答附件
	AnswerAttachmentLink string // TODO
	AnswerAttachmentSize string // TODO
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
	CommentsAttachment     string // 评语附件
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
	// print(stringResponse)
	var homeworkInfo GradedHomeworkInfo

	// Homework Title
	homeworkInfo.Title = strings.Split(stringResponse, "<div class=\"left\">作业标题</div>")[1]
	homeworkInfo.Title = strings.Split(homeworkInfo.Title, "</p>")[0]
	homeworkInfo.Title = stripHtml(homeworkInfo.Title)

	// Homework Description
	homeworkInfo.Description = strings.Split(stringResponse, ">作业说明</div>")[1]
	homeworkInfo.Description = strings.Split(homeworkInfo.Description, "</div>")[0]
	homeworkInfo.Description = stripHtml(homeworkInfo.Description)

	// Homework Attachment
	homeworkInfo.Attachment = strings.Split(stringResponse, ">作业附件</div>")[1]
	homeworkInfo.Attachment = strings.Split(homeworkInfo.Attachment, "</div>")[0]
	homeworkInfo.Attachment = stripHtml(homeworkInfo.Attachment)

	// Homework Answer Description
	homeworkInfo.AnswerDesc = strings.Split(stringResponse, ">答案说明</div>")[1]
	homeworkInfo.AnswerDesc = strings.Split(homeworkInfo.AnswerDesc, "</div>")[0]
	homeworkInfo.AnswerDesc = stripHtml(homeworkInfo.AnswerDesc)

	// Homework Answer Attachment
	homeworkInfo.AnswerAttachment = strings.Split(stringResponse, ">答案附件</div>")[1]
	homeworkInfo.AnswerAttachment = strings.Split(homeworkInfo.AnswerAttachment, "</div>")[0]
	homeworkInfo.AnswerAttachment = stripHtml(homeworkInfo.AnswerAttachment)

	// Homework Publish Object
	homeworkInfo.PublishObject = strings.Split(stringResponse, "<div class=\"left\">发布对象</div>")[1]
	homeworkInfo.PublishObject = strings.Split(homeworkInfo.PublishObject, "</p>")[0]
	homeworkInfo.PublishObject = stripHtml(homeworkInfo.PublishObject)

	// Homework Finish Method
	homeworkInfo.FinishType = strings.Split(stringResponse, "<div class=\"left\">完成方式</div>")[1]
	homeworkInfo.FinishType = strings.Split(homeworkInfo.FinishType, "</p>")[0]
	homeworkInfo.FinishType = stripHtml(homeworkInfo.FinishType)

	// Homework Finish Time
	homeworkInfo.DueDate = strings.Split(stringResponse, "<div class=\"left\">截止日期(GMT+8)</div>")[1]
	homeworkInfo.DueDate = strings.Split(homeworkInfo.DueDate, "</p>")[0]
	homeworkInfo.DueDate = stripHtml(homeworkInfo.DueDate)

	// Homework student id
	homeworkInfo.Studentid = strings.Split(stringResponse, "<div class=\"left\">学号</div>")[1]
	homeworkInfo.Studentid = strings.Split(homeworkInfo.Studentid, "</p>")[0]
	homeworkInfo.Studentid = stripHtml(homeworkInfo.Studentid)

	// Homework submission date
	homeworkInfo.SubmissionDate = strings.Split(stringResponse, "<div class=\"left\">提交日期</div>")[1]
	homeworkInfo.SubmissionDate = strings.Split(homeworkInfo.SubmissionDate, "</p>")[0]
	homeworkInfo.SubmissionDate = stripHtml(homeworkInfo.SubmissionDate)

	// Homework submission content
	homeworkInfo.SubmissionContent = strings.Split(stringResponse, ">上交作业内容</div>")[1]
	homeworkInfo.SubmissionContent = strings.Split(homeworkInfo.SubmissionContent, "</div>")[0]
	homeworkInfo.SubmissionContent = stripHtml(homeworkInfo.SubmissionContent)

	// Homework submission attachment name and link
	homeworkInfo.SubmissionAttachmentName = strings.Split(stringResponse, ">上交作业附件</div>")[1]
	homeworkInfo.SubmissionAttachmentName = strings.Split(homeworkInfo.SubmissionAttachmentName, "</a>")[0]
	homeworkInfo.SubmissionAttachmentLink = strings.Split(homeworkInfo.SubmissionAttachmentName, "<a href=\"")[1]
	homeworkInfo.SubmissionAttachmentLink = strings.Split(homeworkInfo.SubmissionAttachmentLink, "\" target=\"_blank\"")[0]
	homeworkInfo.SubmissionAttachmentName = stripHtml(homeworkInfo.SubmissionAttachmentName)

	// Homework review teacher
	homeworkInfo.ReviewTeacher = strings.Split(stringResponse, "<div class=\"left\">批阅老师</div>")[1]
	homeworkInfo.ReviewTeacher = strings.Split(homeworkInfo.ReviewTeacher, "</p>")[0]
	homeworkInfo.ReviewTeacher = stripHtml(homeworkInfo.ReviewTeacher)

	// Homework review time
	homeworkInfo.ReviewTime = strings.Split(stringResponse, "<div class=\"left\">批阅时间</div>")[1]
	homeworkInfo.ReviewTime = strings.Split(homeworkInfo.ReviewTime, "</p>")[0]
	homeworkInfo.ReviewTime = stripHtml(homeworkInfo.ReviewTime)

	// Homework grade
	homeworkInfo.Grade = strings.Split(stringResponse, "<div class=\"left\">成绩</div>")[1]
	homeworkInfo.Grade = strings.Split(homeworkInfo.Grade, "</p>")[0]
	homeworkInfo.Grade = stripHtml(homeworkInfo.Grade)

	// Homework comments
	homeworkInfo.Comments = strings.Split(stringResponse, ">评语</div>")[1]
	homeworkInfo.Comments = strings.Split(homeworkInfo.Comments, "</span>")[0]
	homeworkInfo.Comments = stripHtml(homeworkInfo.Comments)

	// Homework comments attachment
	homeworkInfo.CommentsAttachment = strings.Split(stringResponse, ">评语附件</div>")[1]
	homeworkInfo.CommentsAttachment = strings.Split(homeworkInfo.CommentsAttachment, "</div>")[0]
	homeworkInfo.CommentsAttachment = stripHtml(homeworkInfo.CommentsAttachment)

	return &homeworkInfo, nil
}
