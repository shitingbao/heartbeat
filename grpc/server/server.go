package server

import (
	"log"

	"github.com/shitingbao/heartbeat/core"
	"github.com/shitingbao/heartbeat/grpc/heart"
)

// HeartServe 外部调用结构体
type HeartServe struct {
	*heart.UnimplementedHeartServerServer
}

func (s *HeartServe) HeartBeat(cli heart.HeartServer_HeartBeatServer) error {
	sid := ""
	for {
		res, err := cli.Recv()
		if err != nil {
			core.UserHub.DeleteData(sid)
			return err
		}
		sid = res.Id
		core.UserHub.PutData(sid)
		log.Println(res.Id)
	}
}
