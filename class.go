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

type ClassInfo struct {
	// Main Information
	Title                string // 课程 Title
	ClassName            string // 课程名称
	MainProfessor        string // 主讲教师
	HekaiProfessor       string // 合开教师
	ClassCredits         string // 课程学分
	ClassHours           string // 课程学时
	ClassScope           string // 课程开放范围
	ClassMaterialsCount  string // 课程文件数
	ClassHomeworkCount   string // 布置作业书
	ClassDiscussionCount string // 讨论贴数

	// Other Information
	ClassDescription               string // 课程简介
	ClassEnglishDescription        string // 英文课程简介
	ClassSchedule                  string // 进度安排
	ClassAssesmentMethod           string // 考核方式
	ClassReferenceMaterials        string // 教材及参考书
	ClassMainReference             string // 主教材
	ClassReferenceBooks            string // 参考书
	ClassProfessor                 string // 授课教师
	ClassSelectionGuidance         string // 选课指导
	ClassPrerequisites             string // 先修要求
	ClassOpenOfficeHour            string // Open Office Hour
	ClassGradingStandard           string // 成绩评定标准
	TeacherTeachingCharacteristics string // 教师教学特色
}

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
func (s *ClassService) GetClassInformation(classID string, courseType string) (*ClassInfo, error) {
	// Create Request Body
	var emptyData url.Values = url.Values{}
	resp, err := s.client.Request(context.Background(), http.MethodPost, addCSRFTokenToUrl(learnCourseInformation(classID, courseType), s.client.csrf), strings.NewReader(emptyData.Encode()))
	if err != nil {
		return nil, err
	}

	stringResponse := DecodeRequestBodyToString(resp)

	var classInfo ClassInfo
	classInfo.Title = strings.ReplaceAll(strings.Split(strings.Split(stringResponse, "<title>")[1], "</title>")[0], "\r ", "")

	// Basic information
	infoTable := strings.Split(strings.Split(stringResponse, "<div class=\"stu_book\">")[1], "<table>")[1]

	// First row
	infoTableRow := strings.Split(infoTable, "<tr>")[1]
	infoTableCol := strings.Split(infoTableRow, "<td>")
	classInfo.MainProfessor = strings.Split(infoTableCol[2], "</td>")[0]
	classInfo.HekaiProfessor = strings.Split(infoTableCol[4], "</td>")[0]

	// Second row
	infoTableRow = strings.Split(infoTable, "<tr>")[2]
	infoTableCol = strings.Split(infoTableRow, "\">")
	classInfo.ClassCredits = strings.Split(infoTableCol[2], "</td>")[0]
	infoTableCol = strings.Split(infoTableRow, "<td>")
	classInfo.ClassHours = strings.Split(infoTableCol[1], "</td>")[0]

	// Third row
	infoTableRow = strings.Split(infoTable, "<tr>")[3]
	infoTableCol = strings.Split(infoTableRow, "<td>")
	classInfo.ClassScope = strings.Split(infoTableCol[2], "</td>")[0]
	classInfo.ClassMaterialsCount = strings.Split(infoTableCol[4], "</td>")[0]

	// Fourth Row
	infoTableRow = strings.Split(infoTable, "<tr>")[4]
	infoTableCol = strings.Split(infoTableRow, "<td>")
	classInfo.ClassHomeworkCount = strings.Split(infoTableCol[2], "</td>")[0]
	classInfo.ClassDiscussionCount = strings.Split(infoTableCol[4], "</td>")[0]

	// Section 2
	sectionTwoInfo := strings.Split(stringResponse, "<div class=\"section2 clearfix\">")

	classInfo.ClassDescription = strings.Split(sectionTwoInfo[1], "<div class=\"cont\">")[1]
	classInfo.ClassEnglishDescription = strings.Split(sectionTwoInfo[2], "<div class=\"cont\">")[1]
	classInfo.ClassSchedule = strings.Split(sectionTwoInfo[3], "<div class=\"cont\">")[1]
	// Deal with <div class=\"section4 section5 ..."
	classInfo.ClassAssesmentMethod = strings.Split(sectionTwoInfo[4], "<div class=\"cont\">")[1]
	classInfo.ClassAssesmentMethod = strings.Split(classInfo.ClassAssesmentMethod, "<div class=\"section4 section5 clearfix\">")[0]

	classInfo.ClassProfessor = strings.Split(sectionTwoInfo[5], "<div class=\"cont\">")[1]
	classInfo.ClassSelectionGuidance = strings.Split(sectionTwoInfo[6], "<div class=\"cont\">")[1]
	classInfo.ClassPrerequisites = strings.Split(sectionTwoInfo[7], "<div class=\"cont\">")[1]
	classInfo.ClassOpenOfficeHour = strings.Split(sectionTwoInfo[8], "<div class=\"cont\">")[1]
	classInfo.ClassGradingStandard = strings.Split(sectionTwoInfo[9], "<div class=\"cont\">")[1]
	classInfo.TeacherTeachingCharacteristics = strings.Split(sectionTwoInfo[10], "<div class=\"cont\">")[1]
	classInfo.TeacherTeachingCharacteristics = strings.Split(classInfo.TeacherTeachingCharacteristics, "</div>")[0]

	// Deal with <div class=\"section4 section5 ..."
	sectionTwoInfo = strings.Split(stringResponse, "<div class=\"section4 section5 clearfix\">")
	// fmt.Println(sectionTwoInfo)
	classInfo.ClassReferenceMaterials = strings.Split(sectionTwoInfo[1], "<p>")[1]
	classInfo.ClassReferenceBooks = strings.Split(sectionTwoInfo[2], "<p>")[1]

	// TODO Remove All Tag Information
	classInfo.ClassDescription = stripHtml(classInfo.ClassDescription)
	classInfo.ClassEnglishDescription = stripHtml(classInfo.ClassEnglishDescription)
	classInfo.ClassSchedule = stripHtml(classInfo.ClassSchedule)
	classInfo.ClassAssesmentMethod = stripHtml(classInfo.ClassAssesmentMethod)

	classInfo.ClassReferenceMaterials = stripHtml(classInfo.ClassReferenceMaterials)
	classInfo.ClassReferenceBooks = stripHtml(classInfo.ClassReferenceBooks)
	classInfo.ClassProfessor = stripHtml(classInfo.ClassProfessor)

	classInfo.ClassSelectionGuidance = stripHtml(classInfo.ClassSelectionGuidance)
	classInfo.ClassPrerequisites = stripHtml(classInfo.ClassPrerequisites)
	classInfo.ClassOpenOfficeHour = stripHtml(classInfo.ClassOpenOfficeHour)
	classInfo.ClassGradingStandard = stripHtml(classInfo.ClassGradingStandard)
	classInfo.TeacherTeachingCharacteristics = stripHtml(classInfo.TeacherTeachingCharacteristics)

	return &classInfo, nil
}

// GetClassNotices : 获取课程公告
func (s *ClassService) GetClassNotices(classID string, courseType string) {

}
