package message

import (
	"strings"
)

type Ctx interface {
	ToUser(user ...string) Ctx
	ToParty(party ...string) Ctx
	ToTag(tag ...string) Ctx
	Safe(b bool) Ctx
	EnableIDTrans(b bool) Ctx
	Send(msg Sender) (*SendResponse, error)
}

type ctx struct {
	tokenFunc func() (string, error)
	*SendParam
}

//New 发送应用消息
func New(t func() (string, error), agentId int) Ctx {
	return &ctx{
		tokenFunc: t,
		SendParam: &SendParam{
			AgentId: agentId,
		},
	}
}

//ToUser 成员ID列表
func (m *ctx) ToUser(user ...string) Ctx {
	m.SendParam.ToUser = strings.Join(user, "|")
	return m
}

//ToParty 部门ID列表
func (m *ctx) ToParty(party ...string) Ctx {
	m.SendParam.ToParty = strings.Join(party, "|")
	return m
}

//ToTag 标签ID列表
func (m *ctx) ToTag(tag ...string) Ctx {
	m.SendParam.ToTag = strings.Join(tag, "|")
	return m
}

//Safe 	表示是否是保密消息
func (m *ctx) Safe(b bool) Ctx {
	if b {
		m.SendParam.Safe = 1
	}
	return m
}

//EnableIDTrans 表示是否开启id转译
func (m *ctx) EnableIDTrans(b bool) Ctx {
	if b {
		m.SendParam.EnableIDTrans = 1
	}
	return m
}

//Send 发送应用消息
func (m *ctx) Send(msg Sender) (*SendResponse, error) {
	t, err := m.tokenFunc()
	if err != nil {
		return nil, err
	}
	return send(t, strings.NewReader(msg.Send(m.SendParam)))
}
