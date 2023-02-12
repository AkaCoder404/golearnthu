package gothulearn

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// UserService: Provides User information service of Learn Account
type UserService service

// GetUserInformation: Get User Information after logging in, such as name, department, type (student/teacher)
func (s *UserService) GetUserInformation() (string, string, string, error) {
	// Create Request Body
	var emptyData url.Values = url.Values{}
	resp, err := s.client.Request(context.Background(), http.MethodPost, learnStudentCourseListPage(), strings.NewReader(emptyData.Encode()))
	if err != nil {
		return "", "", "", err
	}

	// Handle Response
	stringResponse := DecodeRequestBodyToString(resp)
	var labelSplit = strings.Split(stringResponse, "<label>")

	var name string = labelSplit[1]
	name = strings.Split(name, "</label>")[0]
	var studentTypes string = labelSplit[2]
	studentTypes = strings.Split(studentTypes, "</label>")[0]
	var department string = labelSplit[3]
	department = strings.Split(department, "</label>")[0]

	// replace all white spaces
	name = strings.ReplaceAll(name, " ", "")
	studentTypes = strings.ReplaceAll(studentTypes, " ", "")
	studentTypes = strings.ReplaceAll(studentTypes, "r\n", "")
	studentTypes = strings.ReplaceAll(studentTypes, "\t", "")
	department = strings.ReplaceAll(department, " ", "")
	department = strings.ReplaceAll(department, "\r\n", "")
	department = strings.ReplaceAll(department, "\t", "")

	return name, studentTypes, department, nil
}
