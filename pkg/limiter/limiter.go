package limiter

import "errors"

var (
	PoolFullErr = errors.New("pool is full")
)

type Limiter interface {
	Serve()
	Try() error
}
