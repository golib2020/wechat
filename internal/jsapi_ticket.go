package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golib2020/frame/cache"
	"github.com/golib2020/wechat/internal/response"
)

//GetJsapiTicket jsapi ticket
func GetJsapiTicket(appId string, c cache.Cache) func(string) (string, error) {
	return func(token string) (string, error) {
		key := fmt.Sprintf("jsapi.ticket.%s", appId)
		t, err := c.Get(key)
		if err != nil {
			t, err = getJsapiTicket(token)
			if err != nil {
				return "", err
			}
			if err := c.Set(key, t, time.Second*7000); err != nil {
				return "", err
			}
		}
		return t, nil
	}
}

const ticketGetticketApi = `https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=wx_card`

func getJsapiTicket(token string) (string, error) {
	resp, err := http.Get(fmt.Sprintf(ticketGetticketApi, token))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data := new(struct {
		Ticket string `json:"ticket"`
	})
	if err := response.ReadBody(resp.Body, data); err != nil {
		return "", err
	}
	return data.Ticket, nil
}
