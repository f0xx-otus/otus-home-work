package hw04_lru_cache //nolint:golint,stylecheck
import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	sync.Mutex
	capacity int
	queue    *ListView
	dict     map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{capacity: capacity, queue: NewList(), dict: map[Key]*ListItem{}}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.Lock()
	defer l.Unlock()
	if _, ok := l.dict[key]; ok {
		l.queue.Remove(l.dict[key])
		l.dict[key] = l.queue.PushFront(value)
		return true
	}
	if len(l.dict) < l.capacity {
		newListItem := l.queue.PushFront(value)
		l.dict[key] = newListItem
		return false
	}
	for k, v := range l.dict {
		if v.Value == l.queue.Back().Value {
			delete(l.dict, k)
		}
	}
	l.queue.Remove(l.queue.Back())
	l.dict[key] = l.queue.PushFront(value)
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.Lock()
	defer l.Unlock()
	if _, ok := l.dict[key]; ok {
		item := l.dict[key].Value
		l.queue.MoveToFront(l.dict[key])
		l.dict[key] = l.queue.First
		return item, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	for k, v := range l.dict {
		l.queue.Remove(v)
		delete(l.dict, k)
	}
}
