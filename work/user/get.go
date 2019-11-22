package user

import (
	"fmt"
	"net/http"

	"github.com/golib2020/wechat/internal/response"
)

type User struct {
	Userid           string           `json:"userid"`
	Name             string           `json:"name"`
	Department       []int            `json:"department"`
	Order            []int            `json:"order"`
	Position         string           `json:"position"`
	Mobile           string           `json:"mobile"`
	Gender           string           `json:"gender"`
	Email            string           `json:"email"`
	IsLeaderInDept   []int            `json:"is_leader_in_dept"`
	Avatar           string           `json:"avatar"`
	Telephone        string           `json:"telephone"`
	Enable           int              `json:"enable"`
	Alias            string           `json:"alias"`
	Address          string           `json:"address"`
	Extattr          *Extattr         `json:"extattr"`
	Status           int              `json:"status"`
	QrCode           string           `json:"qr_code"`
	ExternalPosition string           `json:"external_position"`
	ExternalProfile  *ExternalProfile `json:"external_profile"`
}
type Text struct {
	Value string `json:"value"`
}
type Web struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}
type Attrs struct {
	Type int    `json:"type"`
	Name string `json:"name"`
	Text Text   `json:"text,omitempty"`
	Web  Web    `json:"web,omitempty"`
}
type Extattr struct {
	Attrs []*Attrs `json:"attrs"`
}
type Miniprogram struct {
	Appid    string `json:"appid"`
	Pagepath string `json:"pagepath"`
	Title    string `json:"title"`
}
type ExternalAttr struct {
	Type        int          `json:"type"`
	Name        string       `json:"name"`
	Text        *Text        `json:"text,omitempty"`
	Web         *Web         `json:"web,omitempty"`
	Miniprogram *Miniprogram `json:"miniprogram,omitempty"`
}
type ExternalProfile struct {
	ExternalCorpName string          `json:"external_corp_name"`
	ExternalAttr     []*ExternalAttr `json:"external_attr"`
}

func get(token, userId string) (*User, error) {
	apiUrl := fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=%s&userid=%s`, token, userId)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	info := new(User)
	if err := response.ReadBody(resp.Body, info); err != nil {
		return nil, err
	}
	return info, nil
}
