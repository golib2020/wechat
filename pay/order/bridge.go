package order

type bridgeConfig struct {
	AppId     string `xml:"appId"`
	TimeStamp string `xml:"timeStamp"`
	NonceStr  string `xml:"nonceStr"`
	Package   string `xml:"package"`
	SignType  string `xml:"signType"`
	PaySign   string `xml:"sign"`
}
