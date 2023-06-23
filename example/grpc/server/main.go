package main

import (
	"log"

	"github.com/shitingbao/heartbeat/grpc/server"
)

func main() {

	h := server.NewGrpcHeart(
		server.WithIsEndless(true),
		server.WithCallback(callback),
	)
	h.Listen()
}

// 接收到心跳后执行的方法
func callback(id string, mes []byte) {
	log.Println("id222222===:", id, "-mes:", string(mes))
}
