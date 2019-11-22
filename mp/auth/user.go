package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/golib2020/wechat/internal/response"
)

type BaseResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
}

type UserinfoResponse struct {
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
	Errcode    int      `json:"errcode"`
	Errmsg     string   `json:"errmsg"`
}

func user(appid, secret, code string) (*UserinfoResponse, error) {
	if code == "" {
		return nil, errors.New("code can't is empty")
	}
	codeApi := `https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code`
	codeApi = fmt.Sprintf(codeApi, appid, secret, code)
	resp, err := http.Get(codeApi)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	base := new(BaseResponse)
	if err := response.ReadBody(resp.Body, base); err != nil {
		return nil, err
	}
	if base.Openid == "" {
		return nil, errors.New("openid can't is empty")
	}
	userinfo := new(UserinfoResponse)
	if base.Scope == "snsapi_userinfo" {
		userinfoUrl := `https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN`
		userinfoUrl = fmt.Sprintf(userinfoUrl, base.AccessToken, base.Openid)
		resp2, err := http.Get(userinfoUrl)
		if err != nil {
			return nil, err
		}
		defer resp2.Body.Close()
		if err := response.ReadBody(resp.Body, userinfo); err != nil {
			return nil, err
		}
		if userinfo.Openid == "" {
			return nil, errors.New("openid can't is empty")
		}
	}
	userinfo.Openid = base.Openid
	return userinfo, nil
}
