package http_client

import (
	"context"

	resty "github.com/go-resty/resty/v2"
)

type RestyRequestModel struct {
	Method              string
	Headers             map[string]string
	Params              map[string]string
	FormData            map[string]string
	Client              *resty.Client
	Ctx                 context.Context
	URL                 string
	Body                interface{}
	FilterResponse      *ResponseFilter
	Files               []*resty.File
	RequestName         string
	AdditionalTags      []string
	RawResponse         *resty.Response
	AdditionalTelemetry *AdditionalTelemetry
}

type ResponseFilter struct {
	Filter func(req *resty.Response) (*resty.Response, error)
}

type AdditionalTelemetry struct {
	Telemetry func(req *resty.Response) (string, string)
}

const (
	InfobipSendEmailRequestName         = "infobip.send_email"
	InfobipSendSMSRequestName           = "infobip.send_sms"
	InfobipVerifyOTPRequestName         = "infobib.verify_otp"
	InfobipSendOTPRequestName           = "infobib.send_otp"
	InfobipUpdateApplicationRequestName = "infobib.update_application"
	SprintasiaLoginOTPRequestName       = "sprintasia.login_otp"
	SprintasiaSendEmailRequestName      = "sprintasia.send_email"
	SprintasiaSendSMSRequestName        = "sprintasia.send_sms"
	SprintasiaVerifyOTPRequestName      = "sprintasia.verify_otp"
	SprintasiaSendOTPRequestName        = "sprintasia.send_otp"
	Vendor8x8SendSMSRequestName         = "sprintasia.send_sms"
	Vendor8x8VerifyOTPRequestName       = "sprintasia.verify_otp"
	Vendor8x8SendOTPRequestName         = "sprintasia.send_otp"
)
