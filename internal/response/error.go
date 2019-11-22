package response

import (
	"errors"
	"fmt"
)

type Checker interface {
	Check() error
}

type Error struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (e *Error) Check() error {
	if e.Errcode != 0 {
		return errors.New(e.String())
	}
	return nil
}

func (e *Error) String() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Errcode, e.Errmsg)
}
