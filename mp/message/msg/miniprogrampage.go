package msg

import (
	"bytes"
	"encoding/json"
)

type MiniPage struct {
	Base
	Miniprogrampage struct {
		Title        string `json:"title"`
		Appid        string `json:"appid"`
		Pagepath     string `json:"pagepath"`
		ThumbMediaID string `json:"thumb_media_id"`
	} `json:"miniprogrampage"`
}

func (m *MiniPage) Staff(openid, kf string) ([]byte, error) {
	m.ToUser = openid
	if kf != "" {
		m.Customservice = &Custom{KfAccount: kf}
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(m)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

//NewMiniPage 创建小程序消息
func NewMiniPage(title, appid, path, thumbMediaId string) *MiniPage {
	m := new(MiniPage)
	m.MsgType = "miniprogrampage"
	m.Miniprogrampage.Title = title
	m.Miniprogrampage.Appid = appid
	m.Miniprogrampage.Pagepath = path
	m.Miniprogrampage.ThumbMediaID = thumbMediaId
	return m
}
