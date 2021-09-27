package session

import "sync"

// Data server
type SessionData struct {
	ID     string
	Data   map[string]interface{}
	rwLock sync.RWMutex // read and write lock, object is Data
	// expire time
}

// Mgr struct
type Mgr struct {
	Session map[string]SessionData
	rwLock  sync.RWMutex
}

func Get(key string) {

}
