package notify

import (
	"encoding/xml"
	"errors"
	"io"

	"github.com/golib2020/wechat/pay/internal"
)

type PaidResponse struct {
	internal.ResponseError
	Appid              string `xml:"appid"`
	MchId              string `xml:"mch_id"`
	DeviceInfo         string `xml:"device_info"`
	NonceStr           string `xml:"nonce_str"`
	Sign               string `xml:"sign"`
	SignType           string `xml:"sign_type"`
	Openid             string `xml:"openid"`
	IsSubscribe        string `xml:"is_subscribe"`
	TradeType          string `xml:"trade_type"`
	BankType           string `xml:"bank_type"`
	TotalFee           string `xml:"total_fee"`
	SettlementTotalFee string `xml:"settlement_total_fee"`
	FeeType            string `xml:"fee_type"`
	CashFee            string `xml:"cash_fee"`
	CashFeeType        string `xml:"cash_fee_type"`
	CouponFee          string `xml:"coupon_fee"`
	CouponCount        string `xml:"coupon_count"`
	TransactionId      string `xml:"transaction_id"`
	OutTradeNo         string `xml:"out_trade_no"`
	Attach             string `xml:"attach"`
	TimeEnd            string `xml:"time_end"`
}

func (n *ctx) paidSign(r io.Reader) (*PaidResponse, error) {
	notify := new(PaidResponse)
	if err := xml.NewDecoder(r).Decode(notify); err != nil {
		return nil, err
	}
	if err := notify.Check(); err != nil {
		return nil, err
	}
	if notify.Sign != internal.SignCheck(n.mchKey, notify) {
		return nil, errors.New("sign error")
	}
	return notify, nil
}
