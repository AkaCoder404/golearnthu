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
	Kxh       string `json:"kxh"`       // 课序号
	Kcm       string `json:"kcm"`       // 课中文名
	Jsh       string `json:"jsh"`       // 教师号
	Kssj      string `json:"kssj"`      // 开始时间
	Jssj      string `json:"jssj"`      // 结束时间
	Kcnr      string `json:"kcnr"`      // 课程内容
	Xf        string `json:"xf"`        // 学分
	Xs        string `json:"xs"`        // 学时
	Cks       string `json:"cks"`       //
	Jxfs      string `json:"jxfs"`      // 教学方式
	Kclx      string `json:"kclx"`      // 课程类型
	Ksfs      string `json:"ksfs"`      // 考试方式
	Lls       int    `json:"lls"`       // 浏览数
	Dys       int    `json:"dys"`       // 答疑数
	Zys       int    `json:"zys"`       // 作业数
	Tjzys     int    `json:"tjzys"`     // 提交作业数
	Zls       int    `json:"zls"`       // 资料数
	Jslls     int    `json:"jslls"`     // 教师浏览数
	Xslls     int    `json:"xslls"`     // 学生浏览数
	Jsdys     int    `json:"jsdys"`     // 教师答疑数
	Xsdys     int    `json:"xsdys"`     // 学生答疑数
	Kfyhlx    string `json:"kfyhlx"`    // 开放用户类型
	Ddys      int    `json:"ddys"`      // 答疑数
	Tlts      int    `json:"ttls"`      // 讨论帖数
	Ggs       int    `json:"ggs"`       // 公告数
	Xtls      int    `json:"xtls"`      // 学堂历史数？
	Jxbjs     int    `json:"jxbjs"`     // 教学班级数？
	Jxkjs     int    `json:"jxkjs"`     // 教学课件数
	Kcfws     int    `json:"kcfws"`     // 课程范围数
	Wpgs      int    `json:"wpgs"`      // 问卷数?
	Xss       int    `json:"xss"`       // 学生数
	Xsbds     int    `json:"xsbds"`     // 学生报到数
	Xzxss     int    `json:"xzxss"`     // 学生在线数
	Tkxss     int    `json:"tkxss"`     // 退课学生数
	Bbh       string `json:"bbh"`       // 版本号
	Czr       string `json:"czr"`       // 创造人
	Czsj      string `json:"czsj"`      // 创造时间
	Bz        string `json:"bz"`        // 备注
	Studentid string `json:"studentid"` // 学生号
	Jsm       string `json:"jsm"`       // 教师名
	Zyzs      int    `json:"zyzs"`      // 作业总数
	Wjzys     int    `json:"wjzys"`     // 未交作业数
	Xggs      int    `json:"xggs"`      // 未读公告数
	Ggzs      int    `json:"ggzs"`      // 公告总数
	Xkjs      int    `json:"xkjs"`      // 未读课件数
	Xskjs     int    `json:"xskjs"`     // 课件总数
	Fqtls     int    `json:"fqtls"`     // 未读讨论帖数
	Cytls     int    `json:"cytls"`     // 已读讨论帖数
	Xsdyzs    int    `json:"xsdyzs"`    // 学生答疑总数
	Yhddys    int    `json:"yhddys"`    // 已回答答疑数
	Ywkcm     string `json:"ywkcm"`     // 英文课程名
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

// ClassListAll : struct just for getting all classes (different from ClassList)
type ClassListAll struct {
	Result string      `json:"success"`
	Msg    string      `json:"msg"`
	Object ClassObject `json:"object"`
}

// ClassObject  : struct for class object under 所有课程
type ClassObject struct {
	ITotalRecords        string     `json:"iTotalRecords"`
	ITotalDisplayRecords string     `json:"iTotalDisplayRecords"`
	IDisplayStart        string     `json:"iDisplayStart"`
	IDisplayLength       string     `json:"iDisplayLength"`
	SEcho                string     `json:"sEcho"`
	SSearch              string     `json:"sSearch"`
	AaData               []ClassAll `json:"aaData"`
}

// ClassAll : struct for class object under 所有课程
type ClassAll struct {
	Xnxq      string `json:"xnxq"`      // 学年学期
	Wlkcid    string `json:"wlkcid"`    // 课程ID
	Kch       string `json:"kch"`       // 课程号
	Kxh       string `json:"kxh"`       // 课序号
	Kcm       string `json:"kcm"`       // 课程名
	Jsh       string `json:"jsh"`       // 教授号
	StudentId string `json:"studentid"` // 学生号
	Jsm       string `json:"jsm"`       // 教授名
	Kcflm     string `json:"kcflm"`     // 课程分类名
	Ggs       int    `json:"ggs"`       // 公告数
	Jxbjs     int    `json:"jxbjs"`     //
	Jxkjs     int    `json:"jxkjs"`     //
	Zys       int    `json:"zys"`       // 作业数
	Kcdys     int    `json:"kcdys"`     // 课程答疑数
	Tlts      int    `json:"tlts"`      // 讨论帖数
	Lls       int    `json:"lls"`       // 浏览数
	Zyzs      int    `json:"zyzs"`      // 作业总数
	Fqtls     int    `json:"fqtls"`     // 发起讨论数
	Cytls     int    `json:"cytls"`     // 参与讨论数
	Xsdys     int    `json:"xsdys"`     // 学生答疑数
	Yxcs      int    `json:"yxcs"`      //
	Bjgcs     int    `json:"bjgcs"`     //
	Kjlls     int    `json:"Kjlls"`     //
	Gglls     int    `json:"gglls"`     // 公告浏览数
	Tls       int    `json:"tls"`       // 讨论数
	Jslx      string `json:"jslx"`      // 教授类型
	Ywkcm     string `json:"ywkcm"`     // 英文课程名
	Id        string `json:"id"`        // id
}

// GetAllClasses : Get a list of all classes
func (c *ClassService) GetAllClasses() (*ClassListAll, error) {
	// Create request
	var payload url.Values = url.Values{}
	payload.Add("aoData", "[{\"name\":\"sEcho\",\"value\":4},{\"name\":\"iColumns\",\"value\":5},{\"name\":\"sColumns\",\"value\":\",,,,\"},{\"name\":\"iDisplayStart\",\"value\":0},{\"name\":\"iDisplayLength\",\"value\":-1},{\"name\":\"mDataProp_0\",\"value\":\"function\"},{\"name\":\"bSortable_0\",\"value\":false},{\"name\":\"mDataProp_1\",\"value\":\"kcm\"},{\"name\":\"bSortable_1\",\"value\":true},{\"name\":\"mDataProp_2\",\"value\":\"jslx\"},{\"name\":\"bSortable_2\",\"value\":true},{\"name\":\"mDataProp_3\",\"value\":\"xnxq\"},{\"name\":\"bSortable_3\",\"value\":true},{\"name\":\"mDataProp_4\",\"value\":\"jsmc\"},{\"name\":\"bSortable_4\",\"value\":true},{\"name\":\"iSortingCols\",\"value\":0}]")
	resp, err := c.client.Request(context.Background(), http.MethodPost, addCSRFTokenToUrl(learnCourseListAll(), c.client.csrf), strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, err
	}

	// Handle Response
	bytesResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response ClassListAll
	json.Unmarshal(bytesResponse, &response)

	log.Printf("There are a total of %d courses", len(response.Object.AaData))
	return &response, nil
}

// TODO convert ClassAll struct to Class struct
