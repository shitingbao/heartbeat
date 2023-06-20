package core

import (
	"sync"
	"time"
)

type HubWithLock struct {
	data map[string]time.Time
	lock *sync.RWMutex
}

func NewDefaultHub() *HubWithLock {
	return &HubWithLock{
		data: make(map[string]time.Time),
		lock: new(sync.RWMutex),
	}
}

func (h *HubWithLock) PutData(key string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.data[key] = time.Now()
}

func (h *HubWithLock) DeleteData(key string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	delete(h.data, key)
}

func (h *HubWithLock) Length() int {
	h.lock.Lock()
	defer h.lock.Unlock()
	return len(h.data)
}

func (h *HubWithLock) GetData(key string) time.Time {
	h.lock.RLock()
	defer h.lock.RUnlock()

	return h.data[key]
}

func (s *HubWithLock) GetAllData() map[string]time.Time {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.data
}
