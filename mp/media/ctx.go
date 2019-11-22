package media

import (
	"io"
)

type Ctx interface {
	Get(mediaId string) (io.Reader, error)
	GetJssdk(mediaId string) (io.Reader, error)
	UploadImage(path string) (*UploadResponse, error)
	UploadThumb(path string) (*UploadResponse, error)
	UploadVideo(path string) (*UploadResponse, error)
	UploadVoice(path string) (*UploadResponse, error)
}

type ctx struct {
	tokenFunc func() (string, error)
}

func New(t func() (string, error)) Ctx {
	return &ctx{tokenFunc: t}
}

//Get 通过media_id获取素材
func (c *ctx) Get(mediaId string) (io.Reader, error) {
	t, err := c.tokenFunc()
	if err != nil {
		return nil, err
	}
	return get(t, mediaId)
}

//GetJssdk 通过media_id获取高清语音素材
func (c *ctx) GetJssdk(mediaId string) (io.Reader, error) {
	t, err := c.tokenFunc()
	if err != nil {
		return nil, err
	}
	return getJssdk(t, mediaId)
}

//UploadImage 图片上传
func (c *ctx) UploadImage(path string) (*UploadResponse, error) {
	t, err := c.tokenFunc()
	if err != nil {
		return nil, err
	}
	return upload(t, "image", path)
}

//UploadThumb 上传缩略图
func (c *ctx) UploadThumb(path string) (*UploadResponse, error) {
	t, err := c.tokenFunc()
	if err != nil {
		return nil, err
	}
	return upload(t, "thumb", path)
}

//UploadVideo 视频上传
func (c *ctx) UploadVideo(path string) (*UploadResponse, error) {
	t, err := c.tokenFunc()
	if err != nil {
		return nil, err
	}
	return upload(t, "video", path)
}

//UploadVoice 语音上传
func (c *ctx) UploadVoice(path string) (*UploadResponse, error) {
	t, err := c.tokenFunc()
	if err != nil {
		return nil, err
	}
	return upload(t, "voice", path)
}
