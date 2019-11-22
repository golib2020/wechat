package media

import (
	"fmt"
	"net/http"

	"github.com/golib2020/wechat/internal"
	"github.com/golib2020/wechat/internal/response"
)

const mediaUploadApi = `https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s`

type UploadResponse struct {
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int    `json:"created_at"`
}

func upload(token, mold, path string) (*UploadResponse, error) {
	request, err := internal.NewMultipartRequest(fmt.Sprintf(mediaUploadApi, token, mold), "media", path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res := new(UploadResponse)
	if err := response.ReadBody(resp.Body, res); err != nil {
		return nil, err
	}
	return res, nil
}
