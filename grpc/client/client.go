package client

import (
	"context"
	"log"
	"time"

	"github.com/shitingbao/heartbeat/grpc/heart"

	"google.golang.org/grpc"

	_ "google.golang.org/grpc/balancer/grpclb"
	"google.golang.org/grpc/credentials/insecure"
)

// const port = "localhost:4399"

// 客户端拦截器
func Clientinterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("method == %s ; req == %v ; rep == %v ; duration == %s ; error == %v\n", method, req, reply, time.Since(start), err)
	return err
}

func DefaultDialOption() []grpc.DialOption {
	opts := []grpc.DialOption{}
	opts = append(opts, grpc.WithUnaryInterceptor(Clientinterceptor))
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials())) // 另一种简单操作
	return opts
}

func startConnect(port string, opts []grpc.DialOption) {
	conn, err := grpc.Dial(port, opts...)
	if err != nil {
		panic(err)
	}
	// defer conn.Close()
	c := heart.NewHeartServerClient(conn) //新建client

	if err := startHeartBeat(c); err != nil {
		log.Println("startHeartBeat:", err)
		return
	}
}

func startHeartBeat(c heart.HeartServerClient) error {
	cli, err := c.HeartBeat(context.Background())
	if err != nil {
		return err
	}
	tm := time.NewTicker(time.Second * 5)
	defer tm.Stop()
	for {
		select {
		case <-tm.C:
			if err := cli.Send(&heart.Heart{
				Id: "1",
			}); err != nil {
				return err
			}
		}
	}

}
