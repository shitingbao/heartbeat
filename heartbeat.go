package heartbeat

type HeartHub interface {
	Listen()
	// Dial(address string, opts ...grpc.DialOption)
	Reboot()
	Endless()
}

func NewHeartHub() {}
