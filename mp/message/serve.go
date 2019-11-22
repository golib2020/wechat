package message

import (
	"encoding/xml"
	"errors"
	"net/http"
	"sort"
	"strings"

	"github.com/golib2020/wechat/internal"
)

type Replyer interface {
	Reply(openid, gh string) ([]byte, error)
}

type RequestBody struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	MsgId        int64
	Content      string `xml:",omitempty"`
	PicUrl       string `xml:",omitempty"`
	MediaId      string `xml:",omitempty"`
	Format       string `xml:",omitempty"`
	Recognition  string `xml:",omitempty"`
	ThumbMediaId string `xml:",omitempty"`
	LocationX    string `xml:",Location_X,omitempty"`
	LocationY    string `xml:",Location_Y,omitempty"`
	Scale        string `xml:",omitempty"`
	Label        string `xml:",omitempty"`
	//Title        string `xml:",omitempty"`
	//Description  string `xml:",omitempty"`
	//Url          string `xml:",omitempty"`
	Event     string `xml:",omitempty"`
	EventKey  string `xml:",omitempty"`
	Ticket    string `xml:",omitempty"`
	Latitude  string `xml:",omitempty"`
	Longitude string `xml:",omitempty"`
	Precision string `xml:",omitempty"`
}

type replyMessage struct {
	Weight  int
	Content Replyer
}

type Server struct {
	nonce             string
	requestBody       *RequestBody
	replyMessageLists []replyMessage
}

//NewServer 开放出来可以自己操作
func NewServe(nonce string, r *http.Request, header func(*Server) error) ([]byte, error) {
	s := &Server{
		nonce: nonce,
	}
	if r.Method == http.MethodGet {
		return s.valid(r)
	} else {
		return s.reply(r, header)
	}
}

//GetRequestBody 获取用户信息
func (s *Server) GetRequestBody() *RequestBody {
	return s.requestBody
}

//push 添加消息
func (s *Server) Push(msg Replyer) *Server {
	s.replyMessageLists = append(s.replyMessageLists, replyMessage{Weight: 0, Content: msg})
	return s
}

//PushWeigh 添加权重消息
func (s *Server) PushWeigh(msg Replyer, weight int) *Server {
	s.replyMessageLists = append(s.replyMessageLists, replyMessage{Weight: weight, Content: msg})
	return s
}

func (s *Server) valid(r *http.Request) ([]byte, error) {
	var signature, timestamp, nonce, echostr string
	signature = r.URL.Query().Get("signature")
	timestamp = r.URL.Query().Get("timestamp")
	nonce = r.URL.Query().Get("nonce")
	echostr = r.URL.Query().Get("echostr")

	list := []string{timestamp, nonce, "123456"}
	sort.Strings(list)
	hashcode := internal.Sha1([]byte(strings.Join(list, "")))
	if !strings.EqualFold(hashcode, signature) {
		return nil, errors.New("验证失败")
	}
	return []byte(echostr), nil
}

func (s *Server) parseRequestBody(r *http.Request) error {
	body := new(RequestBody)
	if err := xml.NewDecoder(r.Body).Decode(body); err != nil {
		return err
	}
	s.requestBody = body
	return nil
}

func (s *Server) reply(r *http.Request, header func(*Server) error) ([]byte, error) {
	//解析数据
	if err := s.parseRequestBody(r); err != nil {
		return nil, err
	}

	//具体操作
	if err := header(s); err != nil {
		return nil, err
	}

	//空消息
	if len(s.replyMessageLists) < 1 {
		return nil, nil
	}
	sort.SliceStable(s.replyMessageLists, func(i, j int) bool {
		return s.replyMessageLists[i].Weight > s.replyMessageLists[j].Weight
	})
	return s.replyMessageLists[0].Content.Reply(s.requestBody.FromUserName, s.requestBody.ToUserName)
}
