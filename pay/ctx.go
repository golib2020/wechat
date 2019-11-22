package pay

import (
	"crypto/tls"

	"github.com/golib2020/wechat/pay/notify"
	"github.com/golib2020/wechat/pay/order"
	"github.com/golib2020/wechat/pay/transfer"
)

type Ctx interface {
	Order() order.Ctx
	Notify() notify.Ctx
	Transfer() transfer.Ctx
}

type ctx struct {
	appId     string
	mchId     string
	mchKey    string
	tlsConfig *tls.Config
}

func New(appId string, opts ...Option) Ctx {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return &ctx{
		appId:     appId,
		mchId:     o.mchId,
		mchKey:    o.mchKey,
		tlsConfig: o.tlsConfig,
	}
}

//Order 下单
func (p *ctx) Order() order.Ctx {
	return order.New(p.appId, p.mchId, p.mchKey, p.tlsConfig)
}

//Notify 通知
func (p *ctx) Notify() notify.Ctx {
	return notify.New(p.mchKey)
}

//Transfer 转账
func (p *ctx) Transfer() transfer.Ctx {
	return transfer.New(p.appId, p.mchId, p.mchKey, p.tlsConfig)
}
