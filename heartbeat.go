package heartbeat

type HeartHub interface {
	Listen()
	Dial()
	Reboot()
	Endless()
}

func NewHeartHub() {}
