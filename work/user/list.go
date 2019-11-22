package user

import (
	"fmt"
	"net/http"

	"github.com/golib2020/wechat/internal/response"
)

type ListResponse struct {
	Userlist []User `json:"userlist"`
}

func list(token string, departmentId, fetchChild int) (*ListResponse, error) {
	apiUrl := fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token=%s&department_id=%d&fetch_child=%d`, token, departmentId, fetchChild)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := new(ListResponse)
	if err := response.ReadBody(resp.Body, data); err != nil {
		return nil, err
	}
	return data, nil
}
