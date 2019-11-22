package open

import "github.com/golib2020/frame/cache"

type options struct {
	cacheAdapter cache.Cache
}

type Option func(*options)

func WithCache(c cache.Cache) Option {
	return func(o *options) {
		o.cacheAdapter = c
	}
}
