package order

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"net/http"

	internal2 "github.com/golib2020/wechat/internal"
	"github.com/golib2020/wechat/pay/internal"
)

type Ctx interface {
	Unify(r *UnifyParam) (*UnifyResponse, error)
	Refund(r *RefundParam) error
}

type ctx struct {
	appid     string
	mchId     string
	mchKey    string
	tlsConfig *tls.Config
}

func New(appid, mchId, mchKey string, tlsConfig *tls.Config) Ctx {
	return &ctx{
		appid:     appid,
		mchId:     mchId,
		mchKey:    mchKey,
		tlsConfig: tlsConfig,
	}
}

//Unify 统一下单
func (o *ctx) Unify(r *UnifyParam) (*UnifyResponse, error) {
	r.Appid = o.appid
	r.MchId = o.mchId
	//数据签名
	r.NonceStr = internal2.RandomStr(10)
	r.Sign = internal.SignCheck(o.mchKey, r)

	bts, err := xml.Marshal(struct {
		XMLName xml.Name `xml:"xml"`
		*UnifyParam
	}{UnifyParam: r})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(`https://api.mch.weixin.qq.com/pay/unifiedorder`, "", bytes.NewReader(bts))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := new(UnifyResponse)
	if err := xml.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err
	}

	if err := res.Check(); err != nil {
		return nil, err
	}
	return res, nil
}

//Refund 退款
func (o *ctx) Refund(r *RefundParam) error {
	r.Appid = o.appid
	r.MchId = o.mchId
	//数据签名
	r.NonceStr = internal2.RandomStr(10)
	r.Sign = internal.SignCheck(o.mchKey, r)

	bts, err := xml.Marshal(struct {
		XMLName xml.Name `xml:"xml"`
		*RefundParam
	}{RefundParam: r})
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, `https://api.mch.weixin.qq.com/secapi/pay/refund`, bytes.NewReader(bts))
	if err != nil {
		return err
	}
	resp, err := internal.ClientTLS(request, o.tlsConfig)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	res := new(internal.ResponseError)
	err = xml.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return err
	}
	return res.Check()
}
