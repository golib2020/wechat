package auth

import (
	"fmt"
	"time"

	"github.com/golib2020/wechat/internal"
)

type WxConfig struct {
	Debug     bool     `json:"debug"`     // 开启调试模式,调用的所有api的返回值会在客户端alert出来，若要查看传入的参数，可以在pc端打开，参数信息会通过log打出，仅在pc端时才会打印。
	AppId     string   `json:"appId"`     // 必填，公众号的唯一标识
	Timestamp int64    `json:"timestamp"` // 必填，生成签名的时间戳
	NonceStr  string   `json:"nonceStr"`  // 必填，生成签名的随机串
	Signature string   `json:"signature"` // 必填，签名
	JsApiList []string `json:"jsApiList"` // 必填，需要使用的JS接口列表
}

//jssdk 获取WxConfig
func jssdk(ticket, appid, url string, debug bool, jsApiList ...string) (*WxConfig, error) {
	timestamp := time.Now().Unix()
	nonceStr := internal.RandomStr(10)
	wxConfig := &WxConfig{
		Debug:     debug,
		AppId:     appid,
		Timestamp: timestamp,
		NonceStr:  nonceStr,
		Signature: getSignature(ticket, nonceStr, timestamp, url),
		JsApiList: jsApiList,
	}
	return wxConfig, nil
}

const signFormat = `jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s`

func getSignature(jsapiTicket, nonceStr string, timestamp int64, url string) string {
	return internal.Sha1([]byte(fmt.Sprintf(signFormat, jsapiTicket, nonceStr, timestamp, url)))
}
