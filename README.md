# gothulearn

This is a Golang API wrapper aimed to provide a program-friendly interface of Learn website of Tsinghua University. It supports [learn2018](https://learn.tsinghua.edu.cn/).

*API documentation in progress*

I made this to learn about the Go Programming Language

**Key Words:** Tsinghua University, API, Learn, 清华大学，网络学堂

## Installation
To use this package, run `go get -u github.com/AkaCoder404/gothulearn`

## Build From Source

## Usage
Here is an example usage case for logging in and getting the courses of this semester.
```go
package main

import (
	"fmt"
	"os"

	"github.com/AkaCoder404/gothulearn"
)

func main() {
	// 创造新的 API Client
	c := gothulearn.NewLearnClient()

	// 登录账号
	err := c.Auth.Login("username", "password")
	if err != nil {
		fmt.Println("Login failed")
	}
	fmt.Println("登入成功")

	// 登录，获取用户基本信息
	name, studentTypes, department, nil := c.User.GetUserInformation()
	if err != nil {
		fmt.Println("User information retrieval failed")
	}
	fmt.Printf("名字: %s\n用户类型: %s\n用户号: %s\n", name, studentTypes, department)

	// 获取本学期
	currentSemesterID, err := c.Class.GetCurrentAndNextSemester()
	currentSemesterID = "2022-2023-1"
	if err != nil {
		fmt.Println("Failed to get current semester")
	}
	fmt.Printf("本学期 %s", currentSemesterID)

	// 获取本学期课程
	list, err := c.Class.GetCourseList(currentSemesterID, "student")
	if err != nil {
		fmt.Println("Failed to get current semester class list")
	}
	for index := 0; index < len(list.ResultList); index++ {
		class := list.ResultList[index]
		fmt.Printf("%s %s\n", class.Kcm, class.Ywkcm)
	}
```


## Change Log
### Completed
- V1.0.0
  - Initial Commit
  - Login

### Working On
- V1.0.1
  - Get current semester 
  - Get list of classes based on semester
  - Get list of all classes
  - Get class information
  - Get unsubmitted/submitted/graded list of a class
  - Get notification list of a class
  - Get file list of a class
  - Get graded homework information **incomplete*

### TODO
- V1.0.2
  - Get unsubmitted/submmited/graded homework information
  - Download file 
