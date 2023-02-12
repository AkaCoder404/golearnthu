package gothulearn

import "net/http"

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
	// Auth  *AuthService
	// User  *UserService
	// Class *ClassService
	// File  *FileService
}

// service: Wrapper for LearnClient
type service struct {
	client *LearnClient
}
