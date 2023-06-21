package client

import (
	"context"
	"log"
	"time"

	"github.com/pborman/uuid"
	"github.com/shitingbao/heartbeat/grpc/heart"

	"google.golang.org/grpc"

	_ "google.golang.org/grpc/balancer/grpclb"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultClientDuration = time.Second * 2
	defaultAddress        = "127.0.0.1:4399"
)

type HeartClient struct {
	ID      string
	Address string
	D       time.Duration
	Opts    []grpc.DialOption
	Cancel  context.CancelFunc
	Ctx     context.Context
	gconn   *grpc.ClientConn
	mes     []byte
}

func clientinterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("method == %s ; req == %v ; rep == %v ; duration == %s ; error == %v\n", method, req, reply, time.Since(start), err)
	return err
}

func defaultDialOption() []grpc.DialOption {
	opts := []grpc.DialOption{}
	opts = append(opts, grpc.WithUnaryInterceptor(clientinterceptor))
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return opts
}

func NewGrpcHeartClient(ctx context.Context, opts ...Option) *HeartClient {
	o := &option{
		Address: defaultAddress,
		D:       defaultClientDuration,
		Opts:    defaultDialOption(),
	}
	for _, opt := range opts {
		opt(o)
	}
	c, cal := context.WithCancel(ctx)
	conn, err := grpc.Dial(o.Address, o.Opts...)
	if err != nil {
		panic(err)
	}

	return &HeartClient{
		ID:     uuid.New(),
		D:      o.D,
		Cancel: cal,
		Ctx:    c,
		gconn:  conn,
	}
}

func (h *HeartClient) Dial() {
	cli := heart.NewHeartServerClient(h.gconn)
	if err := h.startHeartBeat(cli); err != nil {
		log.Println(err)
		return
	}
}

func (h *HeartClient) startHeartBeat(c heart.HeartServerClient) error {
	cli, err := c.HeartBeat(context.Background())
	if err != nil {
		return err
	}
	tm := time.NewTicker(h.D)
	defer tm.Stop()
	for {
		select {
		case <-tm.C:
			if err := cli.Send(&heart.Heart{
				Id:      h.ID,
				Message: h.mes,
			}); err != nil {
				return err
			}
		case <-h.Ctx.Done():
			h.gconn.Close()
			return nil
		}
	}
}

func (h *HeartClient) GetID() string {
	return h.ID
}

// SetMes 设置心跳包额外信息
// 不建议包太大，影响单次包大小
func (h *HeartClient) SetMes(mes []byte) {
	h.mes = mes
}
