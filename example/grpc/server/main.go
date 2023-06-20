package main

import (
	"log"

	"github.com/shitingbao/heartbeat/grpc/server"
)

func main() {

	h := server.NewGrpcHeart(
		server.WithCallback(callback),
	)
	h.Listen()
}

func callback(id string, mes []byte) {
	log.Println("id:", id, "-mes:", string(mes))
}
