package grpc

import (
	"time"

	"github.com/shitingbao/heartbeat"
	"github.com/shitingbao/heartbeat/grpc/server"
)

const (
	defaultHeartDuration = time.Second * 2
	defaultIsEndless     = false
	defaultPort          = ":4399"
)

// GrpcHeart
// Duration   时间间隔
// D          死亡信号
// Port       :4300,有冒号
// createTime 自动保留创建时间
// isEndless  是否无缝重启
// isReConnect  是否重连
type GrpcHeart struct {
	Duration    time.Duration
	D           chan int // 死亡信号
	Port        string
	createTime  time.Time
	isEndless   bool
	isReConnect bool
}

type Option struct {
	Duration    time.Duration
	IsEndless   bool
	Port        string
	IsReConnect bool
}

type GrpcHeartOption func(*Option)

func WithDuration(d time.Duration) GrpcHeartOption {
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

func WithReConnect(isReConncet bool) GrpcHeartOption {
	return func(o *Option) {
		o.IsReConnect = isReConncet
	}
}

func NewGrpcHeart(opts ...GrpcHeartOption) heartbeat.HeartHub {
	o := &Option{
		Duration:  defaultHeartDuration,
		IsEndless: defaultIsEndless,
		Port:      defaultPort,
	}
	for _, opt := range opts {
		opt(o)
	}
	return &GrpcHeart{
		Duration:    o.Duration,
		D:           make(chan int, 1),
		Port:        o.Port,
		createTime:  time.Now(),
		isEndless:   o.IsEndless,
		isReConnect: o.IsReConnect,
	}
}

func (g *GrpcHeart) Listen() {
	server.ServerLoad(g.Port)
}

func (g *GrpcHeart) Dial()    {}
func (g *GrpcHeart) Reboot()  {}
func (g *GrpcHeart) Endless() {}
