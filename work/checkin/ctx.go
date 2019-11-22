package checkin

import (
	"fmt"
	"strings"
)

type Ctx interface {
	GetCheckinData(dataType, startTime, endTime int, user ...string) (*DataResponse, error)
	GetCheckinOption(dateTime int, user ...string) (*OptionResponse, error)
}

type ctx struct {
	tokenFunc func() (string, error)
}

func New(t func() (string, error)) Ctx {
	return &ctx{
		tokenFunc: t,
	}
}

//GetCheckinData 获取打卡数据
func (c *ctx) GetCheckinData(dataType, startTime, endTime int, user ...string) (*DataResponse, error) {
	t, err := c.tokenFunc()
	if err != nil {
		return nil, err
	}
	format := `
	{
	   "opencheckindatatype": %d,
	   "starttime": %d,
	   "endtime": %d,
	   "useridlist": ["%s"]
	}`
	data := fmt.Sprintf(format, dataType, startTime, endTime, strings.Join(user, `", "`))
	return getCheckinData(t, strings.NewReader(data))
}

//GetCheckinOption 获取打卡规则
func (c *ctx) GetCheckinOption(dateTime int, user ...string) (*OptionResponse, error) {
	t, err := c.tokenFunc()
	if err != nil {
		return nil, err
	}
	format := `
	{
		"datetime": %d,
		"useridlist": ["%s"]
	}`
	data := fmt.Sprintf(format, dateTime, strings.Join(user, `", "`))
	return getCheckinOption(t, strings.NewReader(data))
}
