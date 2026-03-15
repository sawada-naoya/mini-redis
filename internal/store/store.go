package store

import "sync"

// Store はkey-valueデータをメモリ上で保持する
type Store struct {
	mu   sync.RWMutex
	data map[string]string
}

func New() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

// Set はkeyにvalueを保存する
func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
}

// Get はkeyに対応するvalueを取得する
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.data[key]
	return value, ok
}

// Del はkeyを削除する。削除できた場合はtrueを返す
func (s *Store) Del(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[key]; !ok {
		return false
	}

	delete(s.data, key)
	return true
}

// Exists はkeyが存在するかを返す
func (s *Store) Exists(key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.data[key]
	return ok
}
