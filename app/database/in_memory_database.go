package database

import (
	"sync"
)

type InMemoryDatabase interface {
	CreateSync(key string, value string)
	GetSync(key string) string
}

type inMemoryDatabase struct {
	database  map[string]string
	writeLock sync.RWMutex
}

func NewInMemoryDatabase() InMemoryDatabase {
	return &inMemoryDatabase{
		database:  map[string]string{},
		writeLock: sync.RWMutex{},
	}
}

func (i *inMemoryDatabase) CreateSync(key string, value string) {
	i.writeLock.Lock()
	defer i.writeLock.Unlock()
	i.database[key] = value
}

func (i *inMemoryDatabase) GetSync(key string) string {
	i.writeLock.RLock()
	defer i.writeLock.RUnlock()
	return i.database[key]
}
