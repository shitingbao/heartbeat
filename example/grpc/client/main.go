package main

import (
	"context"
	"encoding/json"

	"github.com/shitingbao/heartbeat/grpc/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"
)

func main() {
	cli := client.NewGrpcHeartClient(
		context.Background(),
		client.WithGrpcDialOption([]grpc.DialOption{
			grpc.WithTransportCredentials(alts.NewClientCreds(alts.DefaultClientOptions())),
		}))
	b, _ := json.Marshal("hello")
	cli.SetMes(b)
	cli.Dial()
}
