package user

import (
	"fmt"
	"net/http"

	"github.com/golib2020/wechat/internal/response"
)

func delete(token, userId string) error {
	apiUrl := fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/user/delete?access_token=%s&userid=%s`, token, userId)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return response.ReadBody(resp.Body, nil)
}
