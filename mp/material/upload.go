package material

import (
	"fmt"
	"net/http"

	"github.com/golib2020/wechat/internal"
	"github.com/golib2020/wechat/internal/response"
)

const materialAddMaterialApi = `https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%s&type=%s`

type UploadResponse struct {
	MediaId string `json:"media_id"`
	Url     string `json:"url"`
}

func upload(token, mold, path string, params map[string]string) (*UploadResponse, error) {
	request, err := internal.NewMultipartRequest(
		fmt.Sprintf(materialAddMaterialApi, token, mold),
		"media", path, params,
	)
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

//UploadImage 图片上传
func UploadImage(path string) (*material, error) {
	return upload("image", path, nil)
}

//UploadThumb 上传缩略图
func UploadThumb(path string) (*material, error) {
	return upload("thumb", path, nil)
}

//UploadVideo 视频上传
func UploadVideo(path, title, introduction string) (*material, error) {
	params := make(map[string]string)
	params["description"] = fmt.Sprintf(`{"title":"%s", "introduction":"%s"}`, title, introduction)
	return upload("video", path, params)
}

//UploadVoice 语音上传
func UploadVoice(path string) (*material, error) {
	return upload("voice", path, nil)
}
