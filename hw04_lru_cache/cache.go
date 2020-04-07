package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cacher interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    *ListView
	dict     map[Key]*ListItem
}

func NewCache(capacity int) Cacher {
	return &lruCache{capacity: capacity, queue: NewList(), dict: map[Key]*ListItem{}}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
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
		if v.Value == l.queue.Back() {
			delete(l.dict, k)
		}
	}
	l.queue.Remove(l.queue.Back())
	l.dict[key] = l.queue.PushFront(value)
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if _, ok := l.dict[key]; ok {
		item := l.dict[key].Value
		l.queue.Remove(l.dict[key])
		l.dict[key] = l.queue.PushFront(l.dict[key])
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
