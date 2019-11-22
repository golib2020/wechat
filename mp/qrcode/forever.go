package qrcode

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golib2020/wechat/internal/response"
)

const qrcodeCreateApi = `https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s`
const foreverFormat = `{"action_name": "QR_LIMIT_STR_SCENE", "action_info": {"scene": {"scene_str": "%s"}}}`

func forever(token, scene string) (*Response, error) {
	body := strings.NewReader(fmt.Sprintf(foreverFormat, scene))
	resp, err := http.Post(fmt.Sprintf(qrcodeCreateApi, token), `application/json`, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res := new(Response)
	if err := response.ReadBody(resp.Body, res); err != nil {
		return nil, err
	}
	return res, nil
}
