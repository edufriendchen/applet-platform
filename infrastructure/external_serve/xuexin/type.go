package externalAPI

import "github.com/go-resty/resty/v2"

const (
	VerifyStudentStatusURL = "https://www.chsi.com.cn/xlcx/bg.do?vcode=%s&srcid=bgcx"
)

type XueXinProvider struct {
	client *resty.Client
}

type StudentInfo struct {
	Name string
	Sex  string
}
