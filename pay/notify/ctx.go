package notify

import (
	"fmt"
	"net/http"
)

type Ctx interface {
	Refund(w http.ResponseWriter, r *http.Request, notifyHandle func(info *RefundReqInfo) error)
	Paid(w http.ResponseWriter, r *http.Request, notifyHandle func(res *PaidResponse) error)
}

type ctx struct {
	mchKey string
}

//New 通知相关
func New(mchKey string) Ctx {
	return &ctx{mchKey: mchKey}
}

//Refund 退款通知
func (n *ctx) Refund(w http.ResponseWriter, r *http.Request, notifyHandle func(info *RefundReqInfo) error) {
	defer r.Body.Close()
	info, err := n.refundSign(r.Body)
	if err != nil {
		w.Write([]byte(notifyFail(err)))
		return
	}
	if err := notifyHandle(info); err != nil {
		w.Write([]byte(notifyFail(err)))
		return
	}
	w.Write([]byte(notifySuccess()))
}

//Paid 支付通知
func (n *ctx) Paid(w http.ResponseWriter, r *http.Request, notifyHandle func(res *PaidResponse) error) {
	defer r.Body.Close()
	res, err := n.paidSign(r.Body)
	if err != nil {
		w.Write([]byte(notifyFail(err)))
		return
	}
	if err := notifyHandle(res); err != nil {
		w.Write([]byte(notifyFail(err)))
		return
	}
	w.Write([]byte(notifySuccess()))
}

//notifyFail 通知失败
func notifyFail(msg error) string {
	f := `<xml><return_code><![CDATA[FAIL]]></return_code>return_msg><![CDATA[%s]]></return_msg></xml>`
	return fmt.Sprintf(f, msg)
}

//notifySuccess 通知成功
func notifySuccess() string {
	msg := `<xml><return_code><![CDATA[SUCCESS]]></return_code>return_msg><![CDATA[OK]]></return_msg></xml>`
	return msg
}
