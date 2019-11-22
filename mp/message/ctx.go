package message

import (
	"fmt"
	"io"
	"net/http"
)

type Ctx interface {
	Template(jsonReader io.Reader) (int, error)
	Staff(msg Staffer, openid string, kfs ...string) error
	Serve(w http.ResponseWriter, r *http.Request, header func(*Server) error)
}

type ctx struct {
	appid     string
	tokenFunc func() (string, error)
	nonce     string
}

func New(appid, nonce string, t func() (string, error)) Ctx {
	return &ctx{
		appid:     appid,
		tokenFunc: t,
		nonce:     nonce,
	}
}

//Template 模板消息
func (c *ctx) Template(jsonReader io.Reader) (int, error) {
	t, err := c.tokenFunc()
	if err != nil {
		return 0, err
	}
	return template(t, jsonReader)
}

//Staff 客服消息
func (c *ctx) Staff(msg Staffer, openid string, kfs ...string) error {
	t, err := c.tokenFunc()
	if err != nil {
		return err
	}
	kf := ""
	if len(kfs) > 0 && kfs[0] != "" {
		kf = kfs[0]
	}
	return staff(t, openid, kf, msg)
}

//Serve 消息接收与回复
func (c *ctx) Serve(w http.ResponseWriter, r *http.Request, header func(*Server) error) {
	bts, err := NewServe(c.nonce, r, header)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
	w.Write(bts)
}
