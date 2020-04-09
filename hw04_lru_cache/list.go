package hw04_lru_cache //nolint:golint,stylecheck
import (
	"sync"
)

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	sync.Mutex
	Value      interface{}
	Next, Prev *ListItem
}

type ListView struct {
	sync.Mutex
	len         int
	First, Last *ListItem
}

func NewList() *ListView {
	return &ListView{}
}

func (l *ListView) Len() int {
	return l.len
}

func (l *ListView) Front() *ListItem {
	return l.First
}

func (l *ListView) Back() *ListItem {
	return l.Last
}

func (l *ListView) PushFront(data interface{}) *ListItem {
	l.Lock()
	defer l.Unlock()
	var listItem ListItem
	listItem.Value = data
	if l.First == nil {
		l.First = &listItem
		l.Last = &listItem
		listItem.Next = nil
		listItem.Prev = nil
		l.len++
		return &listItem
	}
	oldFirst := l.First
	listItem.Prev = oldFirst
	listItem.Next = oldFirst.Next
	l.First = &listItem
	oldFirst.Next = l.First
	l.len++
	return &listItem
}

func (l *ListView) PushBack(data interface{}) *ListItem {
	var listItem ListItem
	listItem.Value = data
	if l.Last == nil {
		return l.PushFront(data)
	}
	oldLast := l.Last
	listItem.Next = oldLast
	listItem.Prev = oldLast.Prev
	l.Last = &listItem
	oldLast.Prev = l.Last
	l.len++
	return &listItem
}

func (l *ListView) Remove(item *ListItem) {
	l.Lock()
	defer l.Unlock()
	if l.len == 1 {
		l.First.Value = 0
		l.First = nil
		l.Last = nil
		l.len--
		return
	}
	if item.Prev == nil {
		l.Last.Next.Prev = nil
		l.Last = l.Last.Next
		l.len--
		return
	}
	if item.Next == nil {
		l.First.Prev.Next = nil
		l.First = l.First.Prev
		l.len--
		return
	}
	item.Prev.Next = item.Next
	item.Next.Prev = item.Prev
	l.len--
}

func (l *ListView) MoveToFront(item *ListItem) {
	data := item.Value
	l.Remove(item)
	l.PushFront(data)
}
