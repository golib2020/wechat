package message

import (
	"fmt"
)

type textcard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Btntxt      string `json:"btntxt"`
}

func (t *textcard) Send(param *SendParam) string {
	format := `
	{
		"touser": "%s",
		"toparty": "%s",
		"totag": "%s",
		"agentid": %d,
		"enable_id_trans": %d,
		"msgtype": "textcard",
		"textcard": {
			"title": "%s",
			"description": "%s",
			"url": "%s",
			"btntxt": "%s"
		},
	}`
	return fmt.Sprintf(format, param.ToUser, param.ToParty, param.ToTag, param.AgentId, param.EnableIDTrans, t.Title, t.Description, t.URL, t.Btntxt)
}

//TextCard 文本卡片消息
func TextCard(title, desc, url, btnTxt string) *textcard {
	return &textcard{
		Title:       title,
		Description: desc,
		URL:         url,
		Btntxt:      btnTxt,
	}
}
