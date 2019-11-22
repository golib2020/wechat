package mp

import (
	"github.com/golib2020/wechat/internal"
	"github.com/golib2020/wechat/mp/auth"
	"github.com/golib2020/wechat/mp/material"
	"github.com/golib2020/wechat/mp/media"
	"github.com/golib2020/wechat/mp/message"
	"github.com/golib2020/wechat/mp/qrcode"
)

type Ctx interface {
	Auth() auth.Ctx
	Qrcode() qrcode.Ctx
	Media() media.Ctx
	Material() material.Ctx
	Message() message.Ctx
}

type ctx struct {
	appid      string
	secret     string
	tokenFunc  func() (string, error)
	ticketFunc func(token string) (string, error)
	nonce      string
	aesKey     string
}

func New(appid, secret string, opts ...Option) Ctx {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	if o.cacheAdapter == nil {
		panic("cache adapter cannot is nil")
	}

	return &ctx{
		appid:      appid,
		secret:     secret,
		tokenFunc:  internal.GetAccessToken(appid, secret, o.cacheAdapter),
		ticketFunc: internal.GetJsapiTicket(appid, o.cacheAdapter),
	}
}

//Auth 授权认证
func (c *ctx) Auth() auth.Ctx {
	return auth.New(c.appid, c.secret, c.tokenFunc, c.ticketFunc)
}

//Qrcode 生成二维码
func (c *ctx) Qrcode() qrcode.Ctx {
	return qrcode.New(c.tokenFunc)
}

//Media 临时素材管理
func (c *ctx) Media() media.Ctx {
	return media.New(c.tokenFunc)
}

//Material 永久素材管理
func (c *ctx) Material() material.Ctx {
	return material.New(c.tokenFunc)
}

func (c *ctx) Message() message.Ctx {
	return message.New(c.appid, c.nonce, c.tokenFunc)
}
