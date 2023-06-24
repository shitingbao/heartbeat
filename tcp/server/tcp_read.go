package server

type ReadMes struct {
	N   int
	Mes []byte
}

type UpgradeRead interface {
	ReadMessage(b *ReadMes)
}
