package encryptor

import (
	"encoding/json"
	"fmt"
)

type PhoneNumber struct {
	PhoneNumber     string    `json:"phoneNumber"`
	PurePhoneNumber string    `json:"purePhoneNumber"`
	CountryCode     string    `json:"countryCode"`
	Watermark       watermark `json:"watermark"`
}

//getPhoneNumber 获取用户手机号
func getPhoneNumber(appid string, bts []byte) (*PhoneNumber, error) {
	result := new(PhoneNumber)
	if err := json.Unmarshal(bts, result); err != nil {
		return nil, fmt.Errorf("数据解析错误:%s", err)
	}
	if err := result.Watermark.Check(appid); err != nil {
		return nil, err
	}
	return result, nil
}
