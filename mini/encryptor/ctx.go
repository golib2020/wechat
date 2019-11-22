package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

type Ctx interface {
	GetPhoneNumber(session, iv, encryptData string) (*PhoneNumber, error)
	GetUserInfo(session, iv, encryptData string) (*UserInfo, error)
	DecryptData(session, iv, encryptData string) ([]byte, error)
}

type ctx struct {
	appid string
}

//New 消息解密
func New(appid string) *ctx {
	return &ctx{
		appid: appid,
	}
}

//GetPhone 获取用户手机号
func (c *ctx) GetPhoneNumber(session, iv, encryptData string) (*PhoneNumber, error) {
	bts, err := c.DecryptData(session, iv, encryptData)
	if err != nil {
		return nil, err
	}
	return getPhoneNumber(c.appid, bts)
}

//getUserInfo 获取用户信息
func (c *ctx) GetUserInfo(session, iv, encryptData string) (*UserInfo, error) {
	bts, err := c.DecryptData(session, iv, encryptData)
	if err != nil {
		return nil, err
	}
	return getUserInfo(c.appid, bts)
}

//DecryptData 消息解密
func (c *ctx) DecryptData(session, iv, encryptData string) ([]byte, error) {
	return decryptData(encryptData, iv, session)
}

func decryptData(encryptData, iv, key string) ([]byte, error) {
	if len(key) != 24 {
		return nil, fmt.Errorf("aesKey非法")
	}
	if len(iv) != 24 {
		return nil, fmt.Errorf("aesIv非法")
	}
	data, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return nil, fmt.Errorf("encryptData decode error:%s", err)
	}
	keyBts, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, fmt.Errorf("key decode error:%s", err)
	}
	ivBts, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, fmt.Errorf("iv decode error:%s", err)
	}
	dataLen := len(data)
	block, err := aes.NewCipher(keyBts)
	if err != nil {
		return nil, fmt.Errorf("new cipher error:%s", err)
	}
	blockMode := cipher.NewCBCDecrypter(block, ivBts)
	origData := make([]byte, dataLen)
	blockMode.CryptBlocks(origData, data)
	tint := int(origData[dataLen-1])
	if dataLen-tint > 0 {
		return origData[:(dataLen - tint)], nil
	}
	return origData, nil
}

// --------水印---------
type watermark struct {
	Appid     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

func (w *watermark) Check(appId string) error {
	if w.Appid != appId {
		return errors.New("数据不合法")
	}
	return nil
}
