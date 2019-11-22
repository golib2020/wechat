package message

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/golib2020/wechat/internal/response"
)

type Staffer interface {
	Staff(openid, kf string) ([]byte, error)
}

const messageCustomSendApi = `https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s`

func staff(token, openid, kf string, msg Staffer) error {
	bt, err := msg.Staff(openid, kf)
	if err != nil {
		return err
	}
	fmt.Println(string(bt))
	resp, err := http.Post(
		fmt.Sprintf(messageCustomSendApi, token),
		"application/json",
		bytes.NewBuffer(bt),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return response.ReadBody(resp.Body, nil)
}
