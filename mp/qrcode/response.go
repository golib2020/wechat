package qrcode

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

//Response 返回数据
type Response struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"`
	URL           string `json:"url"`
}

const showQrcode = `https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s`

//Url 生成二维码地址
func (r *Response) Url() string {
	return fmt.Sprintf(showQrcode, url.QueryEscape(r.Ticket))
}

//Run 保存到文件
func (r *Response) Save(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	return r.Writer(file)
}

//Run 保存到文件
func (r *Response) Writer(w io.Writer) error {
	resp, err := http.Get(r.Url())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if _, err := io.Copy(w, resp.Body); err != nil {
		return err
	}
	return nil
}
