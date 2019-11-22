# 微信支付

目前仅实现将要用到的接口，剩下的需要慢慢实现

```go
package main

import (
    "fmt"
    "net/http"
    
    "github.com/golib2020/wechat/pay"
    "github.com/golib2020/wechat/pay/notify"
    "github.com/golib2020/wechat/pay/order"
)

func main() {
    
    ctx := pay.New("", pay.WithMCH("", ""), pay.WithTLS("", ""))
    
    //统一下单
    jsapiOrder := order.JSAPI()
    //jsapiOrder.各种参数 = 值
    //jsapiOrder.[] = []
    res, err := ctx.Order().Unify(jsapiOrder)
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
    
    //支付完成异步通知
    http.HandleFunc("/pay/notify/paid", func(w http.ResponseWriter, r *http.Request) {
        ctx.Notify().Paid(w, r, func(info *notify.PaidResponse) error {
            fmt.Println(info)
            return nil
        })
    })
    
    //退款
    refundParam := order.RefundByOutTradeNumber("", "", 100, 100)
    if err = ctx.Order().Refund(refundParam); err != nil {
        panic(err)
    }
    
    //退款异步通知
    http.HandleFunc("/pay/notify/refund", func(w http.ResponseWriter, r *http.Request) {
        ctx.Notify().Refund(w, r, func(info *notify.RefundReqInfo) error {
            fmt.Println(info)
            return nil
        })
    })
}
```