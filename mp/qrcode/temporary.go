package qrcode

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golib2020/wechat/internal/response"
)

const temporaryFormat = `{"expire_seconds": %d, "action_name": "QR_STR_SCENE", "action_info": {"scene": {"scene_str": "%s"}}}`

func temporary(token, scene string, duration time.Duration) (*Response, error) {
	body := strings.NewReader(fmt.Sprintf(temporaryFormat, int(duration.Seconds()), scene))
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
