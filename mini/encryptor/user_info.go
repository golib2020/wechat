package encryptor

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct {
	Openid    string    `json:"openId"`
	NickName  string    `json:"nickName"`
	Gender    int       `json:"gender"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Country   string    `json:"country"`
	AvatarURL string    `json:"avatarUrl"`
	Unionid   string    `json:"unionId"`
	Watermark watermark `json:"watermark"`
}

//getUserInfo 获取用户信息
func getUserInfo(appid string, bts []byte) (*UserInfo, error) {
	result := new(UserInfo)
	if err := json.Unmarshal(bts, result); err != nil {
		return nil, fmt.Errorf("数据解析错误:%s", err)
	}
	if err := result.Watermark.Check(appid); err != nil {
		return nil, err
	}
	return result, nil
}
