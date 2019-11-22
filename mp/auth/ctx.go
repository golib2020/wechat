package auth

type Ctx interface {
	User(code string) (*UserinfoResponse, error)
	Jssdk(url string, debug bool, jsApiList ...string) (*WxConfig, error)
	GetTicket() (string, error)
	GetAccessToken() (string, error)
}

type ctx struct {
	appid      string
	secret     string
	tokenFunc  func() (string, error)
	ticketFunc func(string) (string, error)
}

func New(appid, secret string, token func() (string, error), ticket func(string) (string, error)) Ctx {
	return &ctx{
		appid:      appid,
		secret:     secret,
		tokenFunc:  token,
		ticketFunc: ticket,
	}
}

func (c *ctx) User(code string) (*UserinfoResponse, error) {
	return user(c.appid, c.secret, code)
}

func (c *ctx) Jssdk(url string, debug bool, jsApiList ...string) (*WxConfig, error) {
	t, err := c.GetTicket()
	if err != nil {
		return nil, err
	}
	return jssdk(t, c.appid, url, debug, jsApiList...)
}

func (c *ctx) GetAccessToken() (string, error) {
	return c.tokenFunc()
}

func (c *ctx) GetTicket() (string, error) {
	t, err := c.tokenFunc()
	if err != nil {
		return "", err
	}
	return c.ticketFunc(t)
}
