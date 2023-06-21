package main

import (
	"context"
	"encoding/json"

	"github.com/shitingbao/heartbeat/grpc/client"
)

func main() {
	cli := client.NewGrpcHeartClient(context.Background())
	b, _ := json.Marshal("hello")
	cli.SetMes(b)
	cli.Dial()
}
