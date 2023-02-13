package gothulearn

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	MAX_SIZE int = 20
)

// Request: Sends a request to Learn
func (c *LearnClient) Request(ctx context.Context, method, url string, body io.Reader) (*http.Response, error) {
	// Create the request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// Set Header for Request
	req.Header = c.header

	// Send Request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		// TODO: Decode to an ErrorResponse struct
		log.Println("Request Response status code not 200")
		return nil, fmt.Errorf("Status code not 200")
	}

	// Response successful
	return resp, nil
}

// DecodeRequestBodyToString: Decode a byte response object into a string
func DecodeRequestBodyToString(response *http.Response) string {
	bytesResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return "DecodeRequestBodyToString failed"
	}
	stringResponse := string(bytesResponse)
	return stringResponse
}

// addCSRFTokenToUrl :
func addCSRFTokenToUrl(url string, csrfToken string) string {
	if strings.Contains(url, "?") {
		url += "&_csrf=" + csrfToken
	} else {
		url += "?&_csrf=" + csrfToken
	}
	return url
}

// idLogin: 登录URL
func idLogin() string {
	return "https://id.tsinghua.edu.cn/do/off/ui/auth/login/post/bb5df85216504820be7bba2b0ae1535b/0?/login.do"
}

// idLoginFormData :
func idLoginFormData(username string, password string) url.Values {
	form := url.Values{}
	form.Add("i_user", username)
	form.Add("i_pass", password)
	form.Add("atOnce", "true")
	log.Println(form)
	return form
}

func LearnAuthRom(ticket string) string {
	return LearnPrefix + "/b/j_spring_security_thauth_roaming_entry?ticket=" + ticket
}

// learnLogout : 推出Learn账号
func learnLogout() string {
	return LearnPrefix + "/f/j_spring_security_logout"
}

// learnStudentCourseListPage : 学生课程列表URL
func learnStudentCourseListPage() string {
	return LearnPrefix + "/f/wlxt/index/course/student/"
}

// 学年学期的URL
func learnSemesterList() string {
	return LearnPrefix + "/b/wlxt/kc/v_wlkc_xs_xktjb_coassb/queryxnxq"
}

// 本学期和下学期的URL
func learnCurrentSemester() string {
	return LearnPrefix + "/b/kc/zhjw_v_code_xnxq/getCurrentAndNextSemester"
}

// semester学习的课程
func learnCourseList(semester string, courseType string) string {
	if courseType == "student" {
		return LearnPrefix + "/b/wlxt/kc/v_wlkc_xs_xkb_kcb_extend/student/loadCourseBySemesterId/" + semester
	} else {
		return LearnPrefix + "/b/kc/v_wlkc_kcb/queryAsorCoCourseList/" + semester + "/0"
	}
}

func learnCourseURL(courseID string, courseType string) string {
	return LearnPrefix + "/f/wlxt/index/course/" + courseType + "/course?wlkcid=" + courseID
}

// 课程信息
func learnCourseInformation(courseID string, courseType string) string {
	return LearnPrefix + "/f/wlxt/kc/v_kcxx_jskcxx/" + courseType + "/beforeXskcxx?wlkcid=" + courseID + "&sfgk=0"
}

// 课程时间和地点
func learnCourseTimeLocation(courseID string) string {
	return LearnPrefix + "/b/kc/v_wlkc_xk_sjddb/detail?id=" + courseID
}

func learnTeacherCourseURL(courseID string) string {
	return LearnPrefix + "/f/wlxt/index/course/teacher/course?wlkcid=" + courseID
}

// 课程文件 Tabs
func learnFilePageList(courseID string, courseType string) string {
	return LearnPrefix + "/b/wlxt/kj/wlkc_kjflb/" + courseType + "/pageList?wlkcid=" + courseID
}

// 课程文件 List
func learnFileList(courseID string, filePageID string, courseType string) string {
	return LearnPrefix + "/b/wlxt/kj/wlkc_kjxxb/" + courseType + "/kjxxb/" + courseID + "/" + filePageID
}

// func learnFileList(courseID string, courseType string) string {
// 	if courseType == "student" {
// 		return LearnPrefix + "/b/wlxt/kj/wlkc_kjxxb/student/kjxxbByWlkcidAndSizeForStudent?wlkcid=" + courseID + "&size=" + fmt.Sprint(MAX_SIZE)
// 	} else {
// 		return LearnPrefix + "/b/wlxt/kj/v_kjxxb_wjwjb/teacher/queryByWlkcid?wlkcid=" + courseID + "&size=" + fmt.Sprint(MAX_SIZE)
// 	}
// }

// 下载文件
func learnFileDownload(fileID string, courseType string, courseID string) string {
	if courseType == "student" {
		return LearnPrefix + "/b/wlxt/kj/wlkc_kjxxb/student/downloadFile?sfgk=0&wjid=" + fileID
	} else {
		return LearnPrefix + "/f/wlxt/kj/wlkc_kjxxb/teacher/beforeView?id=" + fileID + "&wlkcid=" + courseID
	}
}

// TODO: 预览文件
func learnFilePreview() string {
	return LearnPrefix
}

// 消息列表
// https://learn.tsinghua.edu.cn/b/wlxt/kcgg/wlkc_ggb/student/pageListXs?_csrf=8262e82e-3eae-4ef4-a64a-7e88ddb8017e

// 消息列表
func learnNotificationList(courseID string, courseType string) string {
	if courseType == "student" {
		return LearnPrefix + "/b/wlxt/kcgg/wlkc_ggb/student/kcggListXs?wlkcid=" + courseID + "&size=" + fmt.Sprint(MAX_SIZE)
	} else {
		return LearnPrefix + "/b/wlxt/kcgg/wlkc_ggb/teacher/kcggList?wlkcid=" + courseID + "&size=" + fmt.Sprint(MAX_SIZE)
	}
}

// TODO: 消息内容
func learnNotificationDetail(courseID string, notificationID string, courseType string) string {
	if courseType == "student" {
		return LearnPrefix
	} else {
		return LearnPrefix
	}
}

// 消息修改
func learnNotificationEdit(courseType string) string {
	return LearnPrefix + "/b/wlxt/kcgg/wlkc_ggb/" + courseType + "/editKcgg"
}

// 未提交作业
func learnUnsubmittedHomeworkList(courseID string) string {
	return LearnPrefix + "/b/wlxt/kczy/zy/student/zyListWj"
}

// 提交未批阅
func learnSubmittedNotGradedHWList(courseID string) string {
	return LearnPrefix + "/b/wlxt/kczy/zy/student/zyListYjwg"
}

// 已经批阅
func learnGradedHomeworkList(courseID string) string {
	return LearnPrefix + "/b/wlxt/kczy/zy/student/zyListYpg"
}
