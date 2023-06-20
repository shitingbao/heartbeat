package server

import (
	"time"

	"github.com/shitingbao/heartbeat/core"
	"google.golang.org/grpc"
)

type Option struct {
	Duration         time.Duration
	IsEndless        bool
	Port             string
	GrpcServerOption []grpc.ServerOption
	UserHub          core.Hub
	Callback         func(string, []byte)
}

type GrpcHeartOption func(*Option)

func WithHeartDuration(d time.Duration) GrpcHeartOption {
	return func(o *Option) {
		o.Duration = d
	}
}

func WithIsEndless(isEndless bool) GrpcHeartOption {
	return func(o *Option) {
		o.IsEndless = isEndless
	}
}

func WithListenPort(port string) GrpcHeartOption {
	return func(o *Option) {
		o.Port = port
	}
}

func WithGrpcOption(opts []grpc.ServerOption) GrpcHeartOption {
	return func(o *Option) {
		o.GrpcServerOption = opts
	}
}

// WithGrpcHub
// You can customize a data storage model to replace the Hub
func WithGrpcHub(h core.Hub) GrpcHeartOption {
	return func(o *Option) {
		o.UserHub = h
	}
}

// WithCallback
// Execute this method when a heartbeat packet is received
func WithCallback(c func(string, []byte)) GrpcHeartOption {
	return func(o *Option) {
		o.Callback = c
	}
}
