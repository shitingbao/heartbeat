package core

import (
	"time"
)

type Hub interface {
	PutData(key string)
	DeleteData(key string)
	Length() int
	GetData(key string) time.Time
	GetAllData() map[string]time.Time
}
