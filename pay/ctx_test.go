package pay

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/golib2020/wechat/pay/notify"
	"github.com/golib2020/wechat/pay/order"
)

func TestNew(t *testing.T) {

	pay := New("", WithMCH("", ""), WithTLS("", ""))

	//统一下单
	jsapiOrder := order.JSAPI()
	//jsapiOrder.各种参数 = 值
	//jsapiOrder.[] = []
	res, err := pay.Order().Unify(jsapiOrder)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	//支付完成异步通知
	http.HandleFunc("/ctx/notify/paid", func(w http.ResponseWriter, r *http.Request) {
		pay.Notify().Paid(w, r, func(info *notify.PaidResponse) error {
			fmt.Println(info)
			return nil
		})
	})

	//退款
	refundParam := order.RefundByOutTradeNumber("", "", 100, 100)
	if err = pay.Order().Refund(refundParam); err != nil {
		panic(err)
	}

	//退款异步通知
	http.HandleFunc("/ctx/notify/refund", func(w http.ResponseWriter, r *http.Request) {
		pay.Notify().Refund(w, r, func(info *notify.RefundReqInfo) error {
			fmt.Println(info)
			return nil
		})
	})

}
