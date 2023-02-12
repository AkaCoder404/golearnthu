package gothulearn

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ClassService service

// SemesterList : A response for getting the current semester
type SemesterList struct {
	Result     map[string]string // Current Semester
	Message    string
	ResultList []Semester `json:"resultList"` // Remaining Semesters
}

type Semester struct {
	Xnxq   string `json:"xnxq"`   // 学年学期
	Xnxqmc string `json:"xnxqmc"` // 学年学期中文名
	Kssj   string `json:"kssj"`   // 开始时间
	Jssj   string `json:"jssj"`   // 结束时间
	Id     string `json:"id"`     // 学期号
}

// ClassList : A response for getting a list of classes
type ClassList struct {
	CurrentUser string  `json:"currentUser"`
	Message     string  `json:"message"`
	ResultList  []Class `json:"resultList"`
}

// Class: struct containing information on Class
type Class struct {
	Xnxq      string `json:"xnxq"`      // 学年学期
	Wlkcid    string `json:"wlkcid"`    // 课程全ID
	Tpid      string `json:"tpid"`      //
	Kch       string `json:"kch"`       // 课程号
	Kxhnumber string `json:"kxhnumber"` // 课程号
	Kxh       string `json:"kxh"`
	Kcm       string `json:"kcm"`  // 课中文名
	Jsh       string `json:"jsh"`  // 教师号
	Kssj      string `json:"kssj"` // 开始时间
	Jssj      string `json:"jssj"` // 结束时间
	Kcnr      string `json:"kcnr"`
	Xf        string `json:"xf"`
	Xs        string `json:"xs"`
	Cks       string `json:"cks"`
	Jxfs      string `json:"jxfs"`
	Kclx      string `json:"kclx"`
	Ksfs      string `json:"ksfs"`
	Lls       int    `json:"lls"`
	Dys       int    `json:"dys"`
	Zys       int    `json:"zys"`
	Tjzys     int    `json:"tjzys"` // 提交作业数
	Zls       int    `json:"zls"`
	Jslls     int    `json:"jslls"`
	Xslls     int    `json:"xslls"`
	Jsdys     int    `json:"jsdys"`
	Xsdys     int    `json:"xsdys"`
	Kfyhlx    string `json:"kfyhlx"`
	Ddys      int    `json:"ddys"`
	Tlts      int    `json:"ttls"`
	Ggs       int    `json:"ggs"`
	Xtls      int    `json:"xtls"`
	Jxbjs     int    `json:"jxbjs"`
	Jxkjs     int    `json:"jxkjs"` // 教学课件数
	Kcfws     int    `json:"kcfws"`
	Wpgs      int    `json:"wpgs"`
	Xss       int    `json:"xss"`
	Xsbds     int    `json:"xsbds"`
	Xzxss     int    `json:"xzxss"`
	Tkxss     int    `json:"tkxss"`
	Bbh       string `json:"bbh"`
	Czr       string `json:"czr"`
	Czsj      string `json:"czsj"`
	Bz        string `json:"bz"`
	Studentid string `json:"studentid"` // 学生号
	Jsm       string `json:"jsm"`       // 教师名
	Zyzs      int    `json:"zyzs"`      // 作业总数
	Wjzys     int    `json:"wjzys"`     // 未交作业数
	Xggs      int    `json:"xggs"`      // 未读公告数
	Ggzs      int    `json:"ggzs"`      // 公告总数
	Xkjs      int    `json:"xkjs"`      // 未读课件数
	Xskjs     int    `json:"xskjs"`     // 课件总数
	Fqtls     int    `json:"fqtls"`
	Cytls     int    `json:"cytls"`
	Xsdyzs    int    `json:"xsdyzs"`
	Yhddys    int    `json:"yhddys"`
	Ywkcm     string `json:"ywkcm"` // 英文课程名
}

// TODO: getClassName
// TODO: getClassNumber
// TODO: ...

// ClassAttributes : Other Information about classes
type ClassAttributes struct {
}

func (s *ClassService) GetCurrentAndNextSemester() (string, error) {
	// Create Request Body
	var emptyData url.Values = url.Values{}
	resp, err := s.client.Request(context.Background(), http.MethodPost, addCSRFTokenToUrl(learnCurrentSemester(), s.client.csrf), strings.NewReader(emptyData.Encode()))
	if err != nil {
		return "", err
	}

	// Read Json Response
	bytesResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var response SemesterList
	json.Unmarshal(bytesResponse, &response)

	CurrentSemesterID = response.Result["id"]
	log.Println("Current Semester ID:", CurrentSemesterID)
	return CurrentSemesterID, nil
}

func (s *ClassService) GetCourseList(semesterID string, courseType string) (*ClassList, error) {
	// Create Request
	timestamp := time.Now().UnixMilli()
	log.Println("Timestamp", timestamp)
	var emptyData url.Values = url.Values{}
	log.Println("URL", addCSRFTokenToUrl(learnCourseList(semesterID, "student"), s.client.csrf))
	resp, err := s.client.Request(context.Background(), http.MethodPost, addCSRFTokenToUrl(learnCourseList(semesterID, "student"), s.client.csrf), strings.NewReader(emptyData.Encode()))

	// Handle Response
	bytesResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response ClassList
	json.Unmarshal(bytesResponse, &response)

	log.Printf("Course List %s, message %s", response.CurrentUser, response.Message)
	log.Printf("There are %d courses in %s semester", len(response.ResultList), semesterID)
	return &response, nil
}

// GetClassInformation : Parse class information page (which is written in HTML) and return parts
func (s *ClassService) GetClassInformation(classID string, courseType string) (string, string, string, string, string, string, string, string, string, error) {
	// Create Request Body
	var emptyData url.Values = url.Values{}
	resp, err := s.client.Request(context.Background(), http.MethodPost, addCSRFTokenToUrl(learnCurrentSemester(), s.client.csrf), strings.NewReader(emptyData.Encode()))
	if err != nil {
		return "", "", "", "", "", "", "", "", "", err
	}

	// TODO: Handle Response
	// stringResponse := DecodeRequestBodyToString(resp)

	if resp == nil {

	}
	return "", "", "", "", "", "", "", "", "", err
}
