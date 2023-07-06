package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/shitingbao/heartbeat/grpc/client"
)

func main() {
	cli := client.NewGrpcHeartClient(
		context.Background(),
		// client.WithGrpcDialOption([]grpc.DialOption{
		// grpc.WithTransportCredentials(alts.NewClientCreds(alts.DefaultClientOptions())),
		// })
	)
	b, err := json.Marshal("hello")
	if err != nil {
		log.Println("json marsh==:", err)
	}
	cli.SetMes(b)
	cli.Dial()
}
