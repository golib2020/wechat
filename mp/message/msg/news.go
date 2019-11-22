package msg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PicUrl      string `json:"picurl"`
	Url         string `json:"url"`
}

//NewArticle 创建文章
func NewArticle(title, desc, pic, url string) *article {
	return &article{
		Title:       title,
		Description: desc,
		PicUrl:      pic,
		Url:         url,
	}
}

type news struct {
	Base
	News struct {
		Articles []*article `json:"articles"`
	} `json:"news"`
}

func (n *news) Staff(openid, kf string) ([]byte, error) {
	n.ToUser = openid
	if kf != "" {
		n.Customservice = &Custom{KfAccount: kf}
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(n)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (n *news) Reply(openid, gh string) ([]byte, error) {
	newsFormat := `<xml><ToUserName><![CDATA[%s]]></ToUserName><FromUserName><![CDATA[%s]]></FromUserName><CreateTime>%d</CreateTime><MsgType><![CDATA[news]]></MsgType><ArticleCount>%d</ArticleCount><Articles>%s</Articles></xml>`
	articleFormat := `<item><Title><![CDATA[%s]]></Title><Description><![CDATA[%s]]></Description><PicUrl><![CDATA[%s]]></PicUrl><Url><![CDATA[%s]]></Url></item>`
	var articles string
	for _, val := range n.News.Articles {
		articles += fmt.Sprintf(articleFormat, val.Title, val.Description, val.PicUrl, val.Url)
	}
	s := fmt.Sprintf(newsFormat, openid, gh, time.Now().Unix(), len(n.News.Articles), articles)
	return []byte(s), nil
}

//Add 单个添加文章
func (n *news) Add(article *article) *news {
	if len(n.News.Articles) >= 8 {
		return n
	}
	n.News.Articles = append(n.News.Articles, article)
	return n
}

//NewNews 创建图文
func NewNews(articles ...*article) *news {
	n := new(news)
	if len(articles) >= 8 {
		n.News.Articles = articles[:8]
	}
	n.MsgType = "news"
	n.News.Articles = articles
	return n
}
