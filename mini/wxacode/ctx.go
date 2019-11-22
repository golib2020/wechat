package wxacode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"net/http"

	"github.com/golib2020/wechat/internal/response"
)

type Ctx interface {
	Get(body io.Reader) (io.Reader, error)
	GetUnlimited(body io.Reader) (io.Reader, error)
	CreateQRCode(body io.Reader) (io.Reader, error)
}

const (
	GetApi          = `https://api.weixin.qq.com/wxa/getwxacode?access_token=%s`
	GetUnlimitedApi = `https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s`
	CreateQRCodeApi = `https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=%s`
)

type ctx struct {
	tokenFunc func() (string, error)
}

//New 小程序码
func New(f func() (string, error)) *ctx {
	return &ctx{
		tokenFunc: f,
	}
}

func (c *ctx) Get(body io.Reader) (io.Reader, error) {
	return c.create(GetApi, body)
}

func (c *ctx) GetUnlimited(body io.Reader) (io.Reader, error) {
	return c.create(GetUnlimitedApi, body)
}

func (c *ctx) CreateQRCode(body io.Reader) (io.Reader, error) {
	return c.create(CreateQRCodeApi, body)
}

func (c *ctx) create(api string, body io.Reader) (io.Reader, error) {
	accessToken, err := c.tokenFunc()
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(fmt.Sprintf(api, accessToken), "application/json", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := new(struct {
		ContentType string          `json:"contentType"`
		Buffer      json.RawMessage `json:"buffer"`
	})
	if err = response.ReadBody(resp.Body, data); err != nil {
		return nil, err
	}
	return bytes.NewReader(data.Buffer), nil
}
