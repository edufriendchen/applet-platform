package externalAPI

import "github.com/go-resty/resty/v2"

type WeiXinProvider struct {
	client *resty.Client
}
