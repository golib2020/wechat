package msg

//Base
type Base struct {
	ToUser        string  `json:"touser"`
	MsgType       string  `json:"msgtype"`
	Customservice *Custom `json:"customservice,omitempty"`
}

type Custom struct {
	KfAccount string `json:"kf_account,omitempty"`
}
