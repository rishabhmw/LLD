package models

import "sync"

type ConcurrentMap struct {
	m  map[interface{}]interface{}
	mu sync.Mutex
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: map[interface{}]interface{}{},
	}
}

func (m *ConcurrentMap) Put(key interface{}, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[key] = value
}

func (m *ConcurrentMap) Get(key interface{}) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.m[key]
	return v, ok
}

func (m *ConcurrentMap) Remove(key interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.m, key)
}

func (m *ConcurrentMap) List() [][]interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.ListUnsafe()
}

func (m *ConcurrentMap) Do(f func()) {
	m.mu.Lock()
	defer m.mu.Unlock()
	f()
}

// PutUnsafe must only be used inside Do
func (m *ConcurrentMap) PutUnsafe(key interface{}, value interface{}) {
	m.m[key] = value
}

// RemoveUnsafe must only be used inside Do
func (m *ConcurrentMap) RemoveUnsafe(key interface{}) {
	delete(m.m, key)
}

// GetUnsafe must only be used inside Do
func (m *ConcurrentMap) GetUnsafe(key interface{}) (any, bool) {
	v, ok := m.m[key]
	return v, ok
}

func (m *ConcurrentMap) ListUnsafe() [][]interface{} {
	result := make([][]interface{}, 0, len(m.m))
	for k, v := range m.m {
		pair := []interface{}{k, v}
		result = append(result, pair)
	}
	return result
}
