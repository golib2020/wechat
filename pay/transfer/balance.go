package transfer


type BalanceParam struct {
	Appid          string `xml:"appid"`
	MchId          string `xml:"mch_id"`
	NonceStr       string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	PartnerTradeNo string `xml:"partner_trade_no"`
	Openid         string `xml:"openid"`
	CheckName      string `xml:"check_name"`
	ReUserName     string `xml:"re_user_name"`
	Amount         int    `xml:"amount"`
	Desc           string `xml:"desc"`
	SpbillCreateIp string `xml:"spbill_create_ip"`
}

//NewTransferBalance 企业付款到用户零钱参数
func NewTransferBalance(partnerTradeNo, openid, desc string, amount int) *BalanceParam {
	param := &BalanceParam{
		PartnerTradeNo: partnerTradeNo,
		Openid:         openid,
		Amount:         amount,
		Desc:           desc,
		SpbillCreateIp: "127.0.0.1",
		CheckName:      "NO_CHECK",
	}
	return param
}
