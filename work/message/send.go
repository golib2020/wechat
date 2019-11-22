package message

import (
	"fmt"
	"io"
	"net/http"

	"github.com/golib2020/wechat/internal/response"
)

type SendParam struct {
	ToUser        string
	ToParty       string
	ToTag         string
	AgentId       int
	Safe          int
	EnableIDTrans int
}

type Sender interface {
	Send(param *SendParam) string
}

type SendResponse struct {
	Invaliduser  string `json:"invaliduser"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
}

//Send 发送应用消息
func send(token string, body io.Reader) (*SendResponse, error) {
	apiUrl := fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s`, token)
	resp, err := http.Post(apiUrl, "application/json", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res := new(SendResponse)
	if err := response.ReadBody(resp.Body, res); err != nil {
		return nil, err
	}
	return res, nil
}
