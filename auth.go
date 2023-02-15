package gothulearn

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// AuthService: Provides Auth services to Learn Account
type AuthService service

// Login: Login to Learn
func (s *AuthService) Login(username string, password string) error {
	// Create required request body (FormData)
	var form url.Values = idLoginFormData(username, password)

	// Login Request
	resp, err := s.client.Request(context.Background(), http.MethodPost, idLogin(), strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	// Check if Login Credentials are correct
	stringResponse := DecodeRequestBodyToString(resp)
	if strings.Contains(stringResponse, "BAD_CREDENTIALS") {
		log.Println("登录失败")
		return errors.New("登录失败, 密码错误")
	} else if strings.Contains(stringResponse, "SUCCESS") {
		log.Println("登录成功")
	}

	// Redirect to Login Page with Ticket
	var redirectPath string = strings.Split(stringResponse, "window.location.replace(\"")[2]
	redirectPath = strings.Split(redirectPath, "\");\n")[0]
	log.Println("Login RedirectPath 1", redirectPath)
	var emptyData = url.Values{}
	resp, err = s.client.Request(context.Background(), http.MethodPost, redirectPath, strings.NewReader(emptyData.Encode()))
	if err != nil {
		return err
	}
	stringResponse = DecodeRequestBodyToString(resp)
	redirectPath = strings.Split(stringResponse, "window.location=\"")[1]
	redirectPath = strings.Split(redirectPath, "\";")[0]
	log.Println("Login Redirect Path 2", redirectPath)

	// Redirect spring security and get csrf token
	resp, err = s.client.Request(context.Background(), http.MethodPost, LearnPrefix+redirectPath, strings.NewReader((emptyData.Encode())))
	if err != nil {
		return err
	}
	stringResponse = DecodeRequestBodyToString(resp)
	var csrfString string = strings.Split(stringResponse, "&_csrf=")[1]
	s.client.csrf = strings.Split(csrfString, "\")")[0]
	log.Println("CSRF_TOKEN", s.client.csrf)
	return nil
}

// Logout : Logout of Learn
func (s *AuthService) Logout() error {
	// Create request

	// handle response
	return nil
}
