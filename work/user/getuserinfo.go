package user

import (
	"fmt"
	"net/http"

	"github.com/golib2020/wechat/internal/response"
)

type UserinfoResponse struct {
	UserId   string `json:"UserId"`
	DeviceId string `json:"DeviceId"`
	OpenId   string `json:"OpenId"`
}

func getUserinfo(token, code string) (*UserinfoResponse, error) {
	apiUrl := fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s`, token, code)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := new(UserinfoResponse)
	if err := response.ReadBody(resp.Body, data); err != nil {
		return nil, err
	}
	return data, nil
}
