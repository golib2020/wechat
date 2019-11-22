package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/golib2020/wechat/internal/response"
)

type Ctx interface {
	GetAccessToken() (string, error)
	Code2Session(code string) (*Session, error)
	GetPaidUnionId(openid string, args ...string) (string, error)
}

type ctx struct {
	appid     string
	secret    string
	tokenFunc func() (string, error)
}

//New 开放接口
func New(appid, secret string, tokenFunc func() (string, error)) Ctx {
	return &ctx{
		appid:     appid,
		secret:    secret,
		tokenFunc: tokenFunc,
	}
}

//Token 获取小程序全局唯一后台接口调用凭据
func (a *ctx) GetAccessToken() (string, error) {
	return a.tokenFunc()
}

//Code2Session 登录凭证校验
func (a *ctx) Code2Session(code string) (*Session, error) {
	apiUrl := fmt.Sprintf(`https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code`, a.appid, a.secret, code)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res := new(Session)
	if err := response.ReadBody(resp.Body, res); err != nil {
		return nil, err
	}
	return res, nil
}

//GetPaidUnionId 用户支付完成后，获取该用户的unionid
func (a *ctx) GetPaidUnionId(openid string, args ...string) (string, error) {
	token, err := a.tokenFunc()
	if err != nil {
		return "", err
	}
	var apiUrl string
	switch len(args) {
	case 1:
		apiUrl = fmt.Sprintf(`https://api.weixin.qq.com/wxa/getpaidunionid?access_token=%s&openid=%s&transaction_id=%s`, token, openid, args[0])
	case 2:
		apiUrl = fmt.Sprintf(`https://api.weixin.qq.com/wxa/getpaidunionid?access_token=%s&openid=%s&mch_id=%s&out_trade_no=%s`, token, openid, args[0], args[1])
	default:
		return "", errors.New("required: transaction_id / mch_id+out_trade_no")
	}
	resp, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data := new(struct {
		Unionid string `json:"unionid"`
	})
	if err := response.ReadBody(resp.Body, data); err != nil {
		return "", err
	}
	return data.Unionid, nil
}
