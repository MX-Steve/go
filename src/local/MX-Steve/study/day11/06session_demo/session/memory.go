package session

import "fmt"

// memory session

// SessionData methods
func (s *SessionData) Get(key string) (value interface{}, err error) {
	// 1. get read lock
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	value, ok := s.Data[key]
	if !ok {
		err = fmt.Errorf("invalid key")
		return
	}
	return
}

// Set set session
func (s *SessionData) Set(key string, value interface{}) {
	// 1. get read lock
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.Data[key] = value
}

// Del del session
func (s *SessionData) Del(key string, value interface{}) {
	// 1. get read lock
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	delete(s.Data, key)
}
