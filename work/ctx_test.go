package work

import (
	"fmt"
	"testing"

	"github.com/golib2020/frame/f"
)

func TestNew(t *testing.T) {
	ctx := New(
		"",
		"",
		WithAgentId(123456),
		WithCache(f.Cache()),
	)
	user, err := ctx.User().Get("snowman")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(user)
}
