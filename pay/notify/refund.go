package notify

import (
	"encoding/base64"
	"encoding/xml"
	"io"

	internal2 "github.com/golib2020/wechat/internal"
	"github.com/golib2020/wechat/pay/internal"
)

type RefundResponse struct {
	internal.ReturnError
	Appid    string `xml:"appid"`
	MchId    string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	ReqInfo  string `xml:"req_info"`
}

type RefundReqInfo struct {
	TransactionId       string `xml:"transaction_id"`
	OutTradeNo          string `xml:"out_trade_no"`
	RefundId            string `xml:"refund_id"`
	OutRefundNo         string `xml:"out_refund_no"`
	TotalFee            int    `xml:"total_fee"`
	SettlementTotalFee  int    `xml:"settlement_total_fee"`
	RefundFee           int    `xml:"refund_fee"`
	SettlementRefundFee int    `xml:"settlement_refund_fee"`
	RefundStatus        string `xml:"refund_status"`
	SuccessTime         string `xml:"success_time"`
	RefundRecvAccout    string `xml:"refund_recv_accout"`
	RefundAccount       string `xml:"refund_account"`
	RefundRequestSource string `xml:"refund_request_source"`
}

func (n *ctx) refundSign(r io.Reader) (*RefundReqInfo, error) {
	notify := new(RefundResponse)
	if err := xml.NewDecoder(r).Decode(notify); err != nil {
		return nil, err
	}
	if err := notify.Check(); err != nil {
		return nil, err
	}
	src, err := base64.StdEncoding.DecodeString(notify.ReqInfo)
	if err != nil {
		return nil, err
	}
	key := internal2.Md5([]byte(n.mchKey))
	dst, err := internal2.AesECBDecrypt(src, []byte(key))
	if err != nil {
		return nil, err
	}
	reqInfo := new(RefundReqInfo)
	if err := xml.Unmarshal(dst, reqInfo); err != nil {
		return nil, err
	}
	return reqInfo, nil
}
