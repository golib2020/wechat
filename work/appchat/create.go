package appchat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golib2020/wechat/internal/response"
)

type CreateParam struct {
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	Userlist []string `json:"userlist"`
	Chatid   string   `json:"chatid`
}

type CreateResponse struct {
	Chatid string `json:"chatid"`
}

func create(token string, param *CreateParam) (*CreateResponse, error) {
	buffer := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buffer).Encode(param); err != nil {
		return nil, err
	}
	apiUrl := fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/appchat/create?access_token=%s`, token)
	resp, err := http.Post(apiUrl, "application/json", buffer)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	info := new(CreateResponse)
	if err := response.ReadBody(resp.Body, info); err != nil {
		return nil, err
	}
	return info, nil
}
