package externalAPI

import (
	"context"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/edufriendchen/applet-platform/common/http_client"
)

func (n *WeiXinProvider) VerifyStudentStatus(ctx context.Context, code string) (*StudentInfo, error) {

	url := fmt.Sprintf(VerifyStudentStatusURL, code)

	resp, err := http_client.Request("POST", url, nil, nil, nil)

	if err != nil {
		hlog.Error(ctx, "[VerifyStudentStatus] error RestyRequest", err)

		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp)))
	if err != nil {
		hlog.Error(ctx, "[VerifyStudentStatus] error NewDocumentFromReader", err)

		return nil, err
	}

	doc.Find(".report-info-item").Each(func(i int, s *goquery.Selection) {
		// 提取label和value的值
		label := s.Find(".label").Text()
		value := s.Find(".value").Text()
		fmt.Printf("%s: %s\n", label, value)
	})

	return nil, nil
}
