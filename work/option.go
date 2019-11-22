package work

import (
	"github.com/golib2020/frame/cache"
)

type options struct {
	cacheAdapter cache.Cache
	agentId      int
}

type Option func(*options)

func WithAgentId(agentId int) Option {
	return func(o *options) {
		o.agentId = agentId
	}
}

func WithCache(c cache.Cache) Option {
	return func(o *options) {
		o.cacheAdapter = c
	}
}
