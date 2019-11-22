package media

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/golib2020/wechat/internal/response"
)

//高清语音素材获取接口地址
const mediaGetJssdkApi = "https://api.weixin.qq.com/cgi-bin/media/get/jssdk?access_token=%s&media_id=%s"

//getJssdk 通过media_id获取高清语音素材
func getJssdk(token, mediaId string) (io.Reader, error) {
	resp, err := http.Get(fmt.Sprintf(mediaGetJssdkApi, token, mediaId))
	if err != nil {
		return nil, fmt.Errorf("http.get error: %s", err)
	}
	defer resp.Body.Close()
	if resp.Header.Get("Content-Type") != "voice/speex" {
		return nil, fmt.Errorf("Content-Type not is voice/speex")
	}
	buf := bytes.NewBuffer(nil)
	if err := response.ReadBody(resp.Body, buf); err != nil {
		return nil, err
	}
	return buf, nil
}
