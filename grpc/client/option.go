package client

import (
	"time"

	"google.golang.org/grpc"
)

type Option func(*option)

type option struct {
	Address string
	D       time.Duration
	Opts    []grpc.DialOption
}

func WithDialAddress(address string) Option {
	return func(o *option) {
		o.Address = address
	}
}

func WithDialDuration(d time.Duration) Option {
	return func(o *option) {
		o.D = d
	}
}

func WithGrpcDialOption(opts []grpc.DialOption) Option {
	return func(o *option) {
		if len(opts) == 0 {
			return
		}
		o.Opts = opts
	}
}
