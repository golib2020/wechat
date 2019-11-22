package qrcode

import "time"

type Ctx interface {
	Forever(scene string) (*Response, error)
	Temporary(scene string, duration time.Duration) (*Response, error)
}

type ctx struct {
	tokenFunc func() (string, error)
}

func New(t func() (string, error)) Ctx {
	return &ctx{tokenFunc: t}
}

//Forever 永久二维码，全部是字符串形式
func (c *ctx) Forever(scene string) (*Response, error) {
	t, err := c.tokenFunc()
	if err != nil {
		return nil, err
	}
	return forever(t, scene)
}

//Temporary 临时二维码，全部是字符串形式
func (c *ctx) Temporary(scene string, duration time.Duration) (*Response, error) {
	t, err := c.tokenFunc()
	if err != nil {
		return nil, err
	}
	return temporary(t, scene, duration)
}
