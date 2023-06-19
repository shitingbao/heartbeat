package core

import (
	"sync"
	"time"
)

type Hub struct {
	data map[string]time.Time
	lock *sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		data: make(map[string]time.Time),
		lock: new(sync.RWMutex),
	}
}

func (h *Hub) PutData(key string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.data[key] = time.Now()
}

func (h *Hub) DeleteData(key string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	delete(h.data, key)
}

func (h *Hub) Length() int {
	h.lock.Lock()
	defer h.lock.Unlock()
	return len(h.data)
}

func (h *Hub) GetData(key string) time.Time {
	h.lock.RLock()
	defer h.lock.RUnlock()

	return h.data[key]
}

func (s *Hub) GetAllData() map[string]time.Time {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.data
}
