package message

import (
	"fmt"
	"strings"
	"time"
)

type news struct {
	Articles []*article `json:"articles"`
}

func (n *news) Send(param *SendParam) string {
	format := `
	{
		"touser": "%s",
		"toparty": "%s",
		"totag": "%s",
		"agentid": %d,
		"safe": %d,
		"enable_id_trans": %d,
		"msgtype": "news",
		"news": {
			"articles": [
				%s
			]
		}
	}`
	articleFormat := `
	{
		"title": "%s",
		"description": "%s",
		"url": "%s",
		"picurl": "%s"
	}`
	var articles []string
	for k, v := range n.Articles {
		if k >= 8 {
			break
		}
		articles = append(articles, fmt.Sprintf(articleFormat, v.Title, v.Description, v.Url, v.PicUrl))
	}
	return fmt.Sprintf(format, param.ToUser, param.ToParty, param.ToTag, param.AgentId, param.Safe, param.EnableIDTrans, strings.Join(articles, ","))
}

func (n *news) Reply(corpID, userId string) string {
	format := `
	<xml>
	   <ToUserName><![CDATA[%s]]></ToUserName>
	   <FromUserName><![CDATA[%s]]></FromUserName>
	   <CreateTime>%d</CreateTime>
	   <MsgType><![CDATA[news]]></MsgType>
	   <Articles>%s</Articles>
	</xml>`
	articleFormat := `
	<item>
		<Title><![CDATA[%s]]></Title> 
		<Description><![CDATA[%s]]></Description>
		<PicUrl><![CDATA[%s]]></PicUrl>
		<Url><![CDATA[%s]]></Url>
	</item>
	`
	var articles string
	for k, v := range n.Articles {
		if k >= 8 {
			break
		}
		articles += fmt.Sprintf(articleFormat, v.Title, v.Description, v.PicUrl, v.Url)
	}
	return fmt.Sprintf(format, userId, corpID, time.Now().Unix(), articles)
}

//Add 添加文章
func (n *news) Add(article *article) *news {
	n.Articles = append(n.Articles, article)
	return n
}


type article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PicUrl      string `json:"picurl"`
	Url         string `json:"url"`
}

//Article 创建文章
func Article(title, desc, pic, url string) *article {
	return &article{
		Title:       title,
		Description: desc,
		PicUrl:      pic,
		Url:         url,
	}
}

//News 图文消息
func News(article ...*article) *news {
	return &news{Articles: article}
}