package message

import (
	"fmt"
	"io"
	"net/http"

	"github.com/golib2020/wechat/internal/response"
)

const messageTemplateSendApi = `https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s`

type SendResponse struct {
	MsgId int `json:"msgid"`
}

//template 发送消息模板
func template(token string, jsonReader io.Reader) (int, error) {
	resp, err := http.Post(
		fmt.Sprintf(messageTemplateSendApi, token),
		"application/json",
		jsonReader,
	)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	res := new(SendResponse)
	if err := response.ReadBody(resp.Body, res); err != nil {
		return 0, err
	}
	return res.MsgId, err
}
