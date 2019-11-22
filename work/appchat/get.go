package appchat

import (
	"fmt"
	"net/http"

	"github.com/golib2020/wechat/internal/response"
)

type GetResponse struct {
	ChatInfo *ChatInfo `json:"chat_info"`
}
type ChatInfo struct {
	Chatid   string   `json:"chatid"`
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	Userlist []string `json:"userlist"`
}

func get(token, chatid string) (*GetResponse, error) {
	apiUrl := fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/appchat/get?access_token=%s&chatid=%s`, token, chatid)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	info := new(GetResponse)
	if err := response.ReadBody(resp.Body, info); err != nil {
		return nil, err
	}
	return info, nil
}
