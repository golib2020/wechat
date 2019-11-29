package internal

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golib2020/frame/cache"
	"github.com/golib2020/wechat/internal/response"
)

//GetAccessToken 接口调用凭证
func GetAccessToken(appId, secret string, c cache.Cache) func() (string, error) {
	return func() (string, error) {
		key := fmt.Sprintf("access_token.%s.%s", appId, secret)
		var t string
		if err := c.Get(key, &t); err != nil {
			t, err = getAccessToken(appId, secret)
			if err != nil {
				return "", err
			}
			if err := c.Set(key, t, time.Second*7000); err != nil {
				return "", err
			}
		}
		return t, nil
	}
}

const (
	wxTokenApi = `https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s`
	qyTokenApi = `https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s`
)

func getAccessToken(appId, secret string) (string, error) {
	var apiUrl string
	if strings.Contains(appId, "wx") {
		apiUrl = fmt.Sprintf(wxTokenApi, appId, secret)
	} else {
		apiUrl = fmt.Sprintf(qyTokenApi, appId, secret)
	}
	resp, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data := new(struct {
		AccessToken string `json:"access_token"`
	})
	if err := response.ReadBody(resp.Body, data); err != nil {
		return "", err
	}
	return data.AccessToken, nil
}
