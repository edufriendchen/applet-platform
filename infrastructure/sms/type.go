package sms

import "context"

type ISMSProvider interface {
	// SendSMS sending sms non otp, returning error
	SendSMS(context context.Context, sms SMS) (string, error)
	// SendOTP sending otp request, returning messagesID, error
	SendOTP(context context.Context, otp OTP) (string, string, error)
	// VerifyOTP verify verification code and given messageID
	VerifyOTP(context context.Context, otp VerifyOTP) error
}

type SMS struct {
	To      string
	Message string
	From    string
}

type OTP struct {
	To   string
	From string
}

type VerifyOTP struct {
	ID   string
	Code string
}
