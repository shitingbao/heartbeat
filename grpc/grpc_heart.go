package grpc

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/shitingbao/heartbeat"
	"github.com/shitingbao/heartbeat/core"
	"github.com/shitingbao/heartbeat/grpc/heart"
	"google.golang.org/grpc"
)

const (
	defaultHeartDuration = time.Second * 2
	defaultIsEndless     = false
	defaultPort          = ":4399"
)

// GrpcHeart
// Duration   time interval
// Dead       death signal
// Port       :4300, with a colon
// createTime automatically retains the creation time
// isEndless  Whether to restart seamlessly
// isReConnect  whether to reconnect
type GrpcHeart struct {
	*heart.UnimplementedHeartServerServer

	Port             string
	Duration         time.Duration
	Dead             chan int
	UserHub          *core.Hub
	isEndless        bool
	isReConnect      bool
	createTime       time.Time
	grpcServerOption []grpc.ServerOption
}

type Option struct {
	Duration         time.Duration
	IsEndless        bool
	Port             string
	IsReConnect      bool
	GrpcServerOption []grpc.ServerOption
}

type GrpcHeartOption func(*Option)

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	m, err := handler(ctx, req)
	end := time.Now()
	log.Println(info.FullMethod, req, start.Format(time.RFC3339), end.Format(time.RFC3339), err)
	return m, err
}

func defaultGrpcOptions() []grpc.ServerOption {
	opts := []grpc.ServerOption{}
	opts = append(opts, grpc.UnaryInterceptor(unaryInterceptor))
	return opts
}

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

func WithGrpcOption(opts []grpc.ServerOption) GrpcHeartOption {
	return func(o *Option) {
		o.GrpcServerOption = opts
	}
}

func NewGrpcHeart(opts ...GrpcHeartOption) heartbeat.HeartHub {
	o := &Option{
		Duration:         defaultHeartDuration,
		IsEndless:        defaultIsEndless,
		Port:             defaultPort,
		GrpcServerOption: defaultGrpcOptions(),
	}
	for _, opt := range opts {
		opt(o)
	}
	return &GrpcHeart{
		Duration:    o.Duration,
		Dead:        make(chan int, 1),
		Port:        o.Port,
		UserHub:     core.NewHub(),
		createTime:  time.Now(),
		isEndless:   o.IsEndless,
		isReConnect: o.IsReConnect,
	}
}

func (g *GrpcHeart) Listen() {
	g.ServerLoad()
}

func (g *GrpcHeart) Reboot()  {}
func (g *GrpcHeart) Endless() {}

func (g *GrpcHeart) ServerLoad() error {
	lis, err := net.Listen("tcp", g.Port)
	if err != nil {
		return err
	}

	s := grpc.NewServer(g.grpcServerOption...)
	heart.RegisterHeartServerServer(s, g)
	return s.Serve(lis)
}

// HeartBeat
func (s *GrpcHeart) HeartBeat(cli heart.HeartServer_HeartBeatServer) error {
	sid := ""

	for {
		res, err := cli.Recv()
		if err != nil {
			s.UserHub.DeleteData(sid)
			return err
		}
		sid = res.Id
		s.UserHub.PutData(sid)
		log.Println(res.Id)
	}
}

func (g *GrpcHeart) GetClientLength() int {
	return g.UserHub.Length()
}

func (g *GrpcHeart) GetAllClient() map[string]time.Time {
	return g.UserHub.GetAllData()
}

func (g *GrpcHeart) GetClientLastTime(key string) time.Time {
	return g.UserHub.GetData(key)
}

// GetClientStatus If the connection is disconnected or exceeds two heartbeats, it is considered dead
func (g *GrpcHeart) GetClientStatus(key string) bool {
	t := g.UserHub.GetData(key)
	if t.IsZero() {
		return false
	}
	return time.Now().Add(g.Duration * -2).Before(t)
}
