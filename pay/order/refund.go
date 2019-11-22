package order

type RefundParam struct {
	Appid         string `xml:"appid"`
	MchId         string `xml:"mch_id"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	SignType      string `xml:"sign_type,omitempty"`
	TransactionId string `xml:"transaction_id,omitempty"`
	OutTradeNo    string `xml:"out_trade_no,omitempty"`
	OutRefundNo   string `xml:"out_refund_no"`
	TotalFee      int    `xml:"total_fee"`
	RefundFee     int    `xml:"refund_fee"`
	RefundFeeType string `xml:"refund_fee_type,omitempty"`
	RefundDesc    string `xml:"refund_desc,omitempty"`
	RefundAccount string `xml:"refund_account,omitempty"`
	NotifyUrl     string `xml:"notify_url,omitempty"`
}

//RefundByOutTradeNumber 根据商户订单号退款
func RefundByOutTradeNumber(transactionId, outRefundNo string, totalFee, refundFee int) *RefundParam {
	param := &RefundParam{
		TransactionId: transactionId,
		OutRefundNo:   outRefundNo,
		TotalFee:      totalFee,
		RefundFee:     refundFee,
	}
	return param
}

//RefundByTransactionId 根据微信订单号退款
func RefundByTransactionId(outTradeNo, outRefundNo string, totalFee, refundFee int) *RefundParam {
	param := &RefundParam{
		OutTradeNo:  outTradeNo,
		OutRefundNo: outRefundNo,
		TotalFee:    totalFee,
		RefundFee:   refundFee,
	}
	return param
}
