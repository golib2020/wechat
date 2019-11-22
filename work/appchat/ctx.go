package appchat

import (
	"strings"
)

type Ctx interface {
	Create(param *CreateParam) (*CreateResponse, error)
	Update(param *UpdateParam) error
	Get(chatid string) (*GetResponse, error)
	Send(chatid string, msg Sender) error
}

type ctx struct {
	tokenFunc func() (string, error)
}

func New(t func() (string, error)) Ctx {
	return &ctx{
		tokenFunc: t,
	}
}

//create 创建群聊会话
func (u *ctx) Create(param *CreateParam) (*CreateResponse, error) {
	t, err := u.tokenFunc()
	if err != nil {
		return nil, err
	}
	return create(t, param)
}

//update 修改群聊会话
func (u *ctx) Update(param *UpdateParam) error {
	t, err := u.tokenFunc()
	if err != nil {
		return err
	}
	return update(t, param)
}

//get 获取群聊会话
func (u *ctx) Get(chatid string) (*GetResponse, error) {
	t, err := u.tokenFunc()
	if err != nil {
		return nil, err
	}
	return get(t, chatid)
}

//send 应用推送消息
func (u *ctx) Send(chatid string, msg Sender) error {
	t, err := u.tokenFunc()
	if err != nil {
		return err
	}
	return send(t, strings.NewReader(msg.ChatSend(chatid)))
}
