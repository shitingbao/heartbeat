package main

import (
	"log"

	"github.com/shitingbao/heartbeat/grpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"
)

func main() {

	h := server.NewGrpcHeart(
		server.WithIsEndless(true),
		server.WithCallback(callback),
		server.WithGrpcOption([]grpc.ServerOption{
			grpc.Creds(alts.NewServerCreds(alts.DefaultServerOptions())),
		}),
	)
	h.Listen()
}

// 接收到心跳后执行的方法
func callback(id string, mes []byte) {
	log.Println("id222222===:", id, "-mes:", string(mes))
}
