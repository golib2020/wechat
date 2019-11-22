package mp

import "github.com/golib2020/frame/cache"

type options struct {
	cacheAdapter cache.Cache
	nonce        string
	aesKey       string
}

type Option func(*options)

//WithCache 设置缓存
func WithCache(c cache.Cache) Option {
	return func(o *options) {
		o.cacheAdapter = c
	}
}

//WithNonce 设置api口令
func WithNonce(s string) Option {
	return func(o *options) {
		o.nonce = s
	}
}

//WithAesKey 设置AES_KEY，暂时不加密，以后再说
func WithAesKey(s string) Option {
	return func(o *options) {
		o.aesKey = s
	}
}
