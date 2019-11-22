package mini

import (
	"github.com/golib2020/wechat/internal"
	"github.com/golib2020/wechat/mini/auth"
	"github.com/golib2020/wechat/mini/encryptor"
	"github.com/golib2020/wechat/mini/wxacode"
)

type Ctx interface {
	Auth() auth.Ctx
	Wxacode() wxacode.Ctx
	Encryptor() encryptor.Ctx
}

type ctx struct {
	appid     string
	secret    string
	tokenFunc func() (string, error)
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
		appid:     appid,
		secret:    secret,
		tokenFunc: internal.GetAccessToken(appid, secret, o.cacheAdapter),
	}
}

//Auth 开放接口
func (c *ctx) Auth() auth.Ctx {
	return auth.New(c.appid, c.secret, c.tokenFunc)
}

//Wxacode 小程序二维码
func (c *ctx) Wxacode() wxacode.Ctx {
	return wxacode.New(c.tokenFunc)
}

//Encryptor 消息解密
func (c *ctx) Encryptor() encryptor.Ctx {
	return encryptor.New(c.appid)
}
