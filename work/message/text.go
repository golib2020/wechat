package message

import (
	"fmt"
	"time"
)

type text struct {
	Content string `json:"content"`
}

func (t *text) Send(param *SendParam) string {
	format := `
	{
		"touser": "%s",
		"toparty": "%s",
		"totag": "%s",
		"agentid": %d,
		"safe": %d,
		"enable_id_trans": %d,
		"msgtype": "textcard",
		"textcard": {
			"content": "%s"
		}
	}`
	return fmt.Sprintf(format, param.ToUser, param.ToParty, param.ToTag, param.AgentId, param.Safe, param.EnableIDTrans, t.Content)
}

func (t *text) Reply(corpID, userId string) string {
	format := `
	<xml>
	   <ToUserName><![CDATA[%s]]></ToUserName>
	   <FromUserName><![CDATA[%s]]></FromUserName>
	   <CreateTime>%d</CreateTime>
	   <MsgType><![CDATA[textcard]]></MsgType>
	   <Content><![CDATA[%s]]></Content>
	</xml>`
	return fmt.Sprintf(format, userId, corpID, time.Now().Unix(), t.Content)
}

func (t *text) ChatSend(chatid string) string {
	format := `
	{
		"chatid": "%s",
		"msgtype":"text",
		"text":{
			"content" : "%s"
		},
		"safe":0
	}`
	return fmt.Sprintf(format, chatid, t.Content)
}

//Text 文本消息
func Text(content string) *text {
	return &text{Content: content}
}
