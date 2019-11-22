package work

import (
	"github.com/golib2020/wechat/internal"
	"github.com/golib2020/wechat/work/appchat"
	"github.com/golib2020/wechat/work/checkin"
	"github.com/golib2020/wechat/work/message"
	"github.com/golib2020/wechat/work/user"
)

type Ctx interface {
	User() user.Ctx
	Message() message.Ctx
	Checkin() checkin.Ctx
	AppChat() appchat.Ctx
}

type ctx struct {
	corpId     string
	corpSecret string
	tokenFunc  func() (string, error)
	agentId    int
}

func New(appId, secret string, opts ...Option) Ctx {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	if o.cacheAdapter == nil {
		panic("cache adapter cannot is nil")
	}
	return &ctx{
		corpId:     appId,
		corpSecret: secret,
		tokenFunc:  internal.GetAccessToken(appId, secret, o.cacheAdapter),
		agentId:    o.agentId,
	}
}

//ctx 成员相关
func (a *ctx) User() user.Ctx {
	return user.New(a.tokenFunc)
}

//ctx 消息相关
func (a *ctx) Message() message.Ctx {
	return message.New(a.tokenFunc, a.agentId)
}

//ctx 打卡相关
func (a *ctx) Checkin() checkin.Ctx {
	return checkin.New(a.tokenFunc)
}

//ctx 群聊相关
func (a *ctx) AppChat() appchat.Ctx {
	return appchat.New(a.tokenFunc)
}
