package main

import (
	"context"

	"github.com/shitingbao/heartbeat/grpc/client"
)

func main() {
	cli := client.NewGrpcHeartClient(context.Background())
	cli.Dial()
}
