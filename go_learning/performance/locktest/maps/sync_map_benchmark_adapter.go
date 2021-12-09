package maps

import "sync"

// sync.map相关接口实现
func CreateSyncMapBenchmarkAdapter() *SyncMapBenchmarkAdapter {
	return &SyncMapBenchmarkAdapter{}
}

type SyncMapBenchmarkAdapter struct {
	m sync.Map
}

func (m *SyncMapBenchmarkAdapter) Set(key interface{}, val interface{}) {
	m.m.Store(key, val)
}

func (m *SyncMapBenchmarkAdapter) Get(key interface{}) (interface{}, bool) {
	return m.m.Load(key)
}

func (m *SyncMapBenchmarkAdapter) Del(key interface{}) {
	m.m.Delete(key)
}
