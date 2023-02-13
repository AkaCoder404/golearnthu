package gothulearn

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
)

// Global variables
var (
	CurrentSemesterID string
	NextSemesterID    string
	CourseType        string
)

// Global Consts
const (
	LearnPrefix     = "https://learn.tsinghua.edu.cn"
	RegistrarPrefix = "https://zhjw.cic.tsinghua.edu.cn"
)

// LearnClient: The GoLearn client
type LearnClient struct {
	client *http.Client
	header http.Header

	common       service
	refreshToken string
	csrf         string

	// Services for LearnClient API
	Auth     *AuthService
	User     *UserService
	Class    *ClassService
	File     *FileService
	Homework *HomeworkService
	Notice   *NoticeService
}

// service: Wrapper for LearnClient
type service struct {
	client *LearnClient
}

// NewLearnClient: New anonymous client. To login as authenticated user, use LearnClient.Login
func NewLearnClient() *LearnClient {
	// Create client
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Error Creating NewLearn Client")
	}
	client := http.Client{Jar: jar}

	// Create header
	header := http.Header{}
	header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create the new client
	learn := &LearnClient{
		client: &client,
		header: header,
	}

	// TODO: Set the common client
	learn.common.client = learn

	// Reuse the common client for the other services
	learn.Auth = (*AuthService)(&learn.common)
	learn.User = (*UserService)(&learn.common)
	learn.Class = (*ClassService)(&learn.common)
	learn.File = (*FileService)(&learn.common)
	learn.Homework = (*HomeworkService)(&learn.common)
	learn.Notice = (*NoticeService)(&learn.common)
	return learn
}
