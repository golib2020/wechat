package message

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golib2020/wechat/mp/message/msg"
)

func home(w http.ResponseWriter, r *http.Request) {
	//reply serve
	bts, err := NewServe("123456", r, func(s *Server) error {
		s.GetRequestBody()
		s.Push(msg.NewText("111"))
		return nil
	})
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
	w.Write(bts)
}

func TestNewServe(t *testing.T) {

	http.HandleFunc("/", home)
	ts := httptest.NewServer(nil)


	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s", bts)

}
