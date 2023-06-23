package server

import (
	"context"
	"log"
	"net"
	"sync"
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
// Port       example ":4300", with a colon
// Duration   default heart time interval
// Dead       death signal
// UserHub    data hub
// listen     useing net.TCPListener
// wg		  Used to ensure that the old process does not stop until all connections of the old process exit
// createTime automatically retains the creation time
// isEndless  Whether to restart seamlessly
type GrpcHeart struct {
	*heart.UnimplementedHeartServerServer

	Port     string
	Duration time.Duration
	Dead     chan int
	UserHub  core.Hub

	listen net.Listener
	wg     *sync.WaitGroup

	isEndless        bool
	createTime       time.Time
	grpcServerOption []grpc.ServerOption
	callback         func(string, []byte)
}

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

// NewGrpcHeart
func NewGrpcHeart(opts ...GrpcHeartOption) heartbeat.HeartHub {
	o := &Option{
		Duration:         defaultHeartDuration,
		IsEndless:        defaultIsEndless,
		Port:             defaultPort,
		GrpcServerOption: defaultGrpcOptions(),
		UserHub:          core.NewDefaultHub(),
		Callback:         func(string, []byte) {},
	}
	for _, opt := range opts {
		opt(o)
	}
	return &GrpcHeart{
		Port:             o.Port,
		Duration:         o.Duration,
		Dead:             make(chan int, 1),
		UserHub:          o.UserHub,
		wg:               &sync.WaitGroup{},
		createTime:       time.Now(),
		grpcServerOption: o.GrpcServerOption,
		isEndless:        o.IsEndless,
		callback:         o.Callback,
	}
}

func (g *GrpcHeart) Listen() {
	if g.isEndless {
		g.endlessTcpRegisterAndListen()
	} else {
		lis, err := net.Listen("tcp", g.Port)
		if err != nil {
			panic(err)
		}
		g.listen = lis
		g.serverLoad()
	}
}

func (g *GrpcHeart) serverLoad() error {
	s := grpc.NewServer(g.grpcServerOption...)
	heart.RegisterHeartServerServer(s, g)
	log.Println("start listen:", g.Port)
	return s.Serve(g.listen)
}

// HeartBeat
// Implemented grpc method
func (g *GrpcHeart) HeartBeat(cli heart.HeartServer_HeartBeatServer) error {
	sid := ""
	g.wg.Add(1)
	defer func() {
		g.wg.Done()

	}()
	for {
		res, err := cli.Recv()
		if err != nil {
			g.UserHub.DeleteData(sid)
			return err
		}
		g.callback(res.Id, res.Message)
		sid = res.Id
		g.UserHub.PutData(sid)
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
