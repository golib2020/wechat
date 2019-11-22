package pay

import (
	"crypto/tls"
	"log"
)

type options struct {
	mchId     string
	mchKey    string
	tlsConfig *tls.Config
}

type Option func(*options)

//WithMCH 商户帐号
func WithMCH(mchId, mchKey string) Option {
	return func(opts *options) {
		opts.mchId = mchId
		opts.mchKey = mchKey
	}
}

//WithTLS 证书
func WithTLS(certPath, keyPath string) Option {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Panic("cert load fail:", err)
	}
	return func(opts *options) {
		opts.tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
	}
}
