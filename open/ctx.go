package open

import (
	"github.com/golib2020/wechat/internal"
)

type ctx struct {
	appid     string
	secret    string
	tokenFunc func() (string, error)
}

func New(appid, secret string, opts ...Option) *ctx {
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
