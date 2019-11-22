package msg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type text struct {
	Base
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func (t *text) Staff(openid, kf string) ([]byte, error) {
	t.ToUser = openid
	if kf != "" {
		t.Customservice = &Custom{KfAccount: kf}
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (t *text) Reply(openid, gh string) ([]byte, error) {
	textFormat := `<xml><ToUserName><![CDATA[%s]]></ToUserName><FromUserName><![CDATA[%s]]></FromUserName><CreateTime>%d</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[%s]]></Content></xml>`
	s := fmt.Sprintf(textFormat, openid, gh, time.Now().Unix(), t.Text.Content)
	return []byte(s), nil
}

func NewText(content string) *text {
	t := new(text)
	t.MsgType = "text"
	t.Text.Content = content
	return t
}
