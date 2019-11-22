package transfer

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"net/http"

	internal2 "github.com/golib2020/wechat/internal"
	"github.com/golib2020/wechat/pay/internal"
)

type Ctx interface {
	Balance(r *BalanceParam) error
}

type ctx struct {
	appid     string
	mchId     string
	mchKey    string
	tlsConfig *tls.Config
}

//New 企业转账
func New(appid, mchId, mchKey string, tlsConfig *tls.Config) Ctx {
	return &ctx{
		appid:     appid,
		mchId:     mchId,
		mchKey:    mchKey,
		tlsConfig: tlsConfig,
	}
}

//Balance 企业付款到用户零钱
func (t *ctx) Balance(r *BalanceParam) error {
	r.Appid = t.appid
	r.MchId = t.mchId
	//数据签名
	r.NonceStr = internal2.RandomStr(10)
	r.Sign = internal.SignCheck(t.mchKey, r)

	bts, err := xml.Marshal(struct {
		XMLName xml.Name `xml:"xml"`
		*BalanceParam
	}{BalanceParam: r})
	if err != nil {
		return err
	}
	//fmt.Println(string(bts))
	request, err := http.NewRequest(http.MethodPost, `https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers`, bytes.NewReader(bts))
	if err != nil {
		return err
	}
	resp, err := internal.ClientTLS(request, t.tlsConfig)
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
