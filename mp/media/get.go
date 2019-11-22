package media

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/golib2020/wechat/internal/response"
)

const mediaGetApi = "http://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"

//get 通过media_id获取素材
func get(token, mediaId string) (io.Reader, error) {
	resp, err := http.Get(fmt.Sprintf(mediaGetApi, token, mediaId))
	if err != nil {
		return nil, fmt.Errorf("http.get error: %s", err)
	}
	defer resp.Body.Close()
	buf := bytes.NewBuffer(nil)
	if err := response.ReadBody(resp.Body, buf); err != nil {
		return nil, err
	}
	return buf, nil
}
