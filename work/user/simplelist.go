package user

import (
	"fmt"
	"net/http"

	"github.com/golib2020/wechat/internal/response"
)

type SimpleListResponse struct {
	Userlist []Userlist `json:"userlist"`
}

type Userlist struct {
	Userid     string `json:"userid"`
	Name       string `json:"name"`
	Department []int  `json:"department"`
}

func simpleList(token string, departmentId, fetchChild int) (*SimpleListResponse, error) {
	apiUrl := fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=%s&department_id=%d&fetch_child=%d`, token, departmentId, fetchChild)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := new(SimpleListResponse)
	if err := response.ReadBody(resp.Body, data); err != nil {
		return nil, err
	}
	return data, nil
}
