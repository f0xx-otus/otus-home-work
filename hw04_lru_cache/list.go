package hw04_lru_cache //nolint:golint,stylecheck

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
	Value      interface{}
	Next, Prev *ListItem
}

type ListView struct {
	len         int
	first, last *ListItem
}

func NewList() *ListView {
	return &ListView{}
}

func (l *ListView) Len() int {
	return l.len
}

func (l *ListView) Front() *ListItem {
	return l.first
}

func (l *ListView) Back() *ListItem {
	return l.last
}

func (l *ListView) PushFront(data interface{}) *ListItem {
	var listItem ListItem
	listItem.Value = data
	if l.first == nil {
		l.first = &listItem
		l.last = &listItem
		listItem.Next = nil
		listItem.Prev = nil
		l.len++
		return &listItem
	}
	oldFirst := l.first
	listItem.Prev = oldFirst
	listItem.Next = oldFirst.Next
	l.first = &listItem
	oldFirst.Next = l.first
	l.len++
	return &listItem
}

func (l *ListView) PushBack(data interface{}) *ListItem {
	var listItem ListItem
	listItem.Value = data
	if l.last == nil {
		return l.PushFront(data)
	}
	oldLast := l.last
	listItem.Next = oldLast
	listItem.Prev = oldLast.Prev
	l.last = &listItem
	oldLast.Prev = l.last
	l.len++
	return &listItem
}

func (l *ListView) Remove(item *ListItem) {
	if l.len == 1 {
		l.first.Value = 0
		l.first = nil
		l.last = nil
		l.len--
		return
	}
	if item.Prev == nil {
		l.last = l.last.Next
		l.last.Prev = nil
		l.len--
		return
	}
	if item.Next == nil {
		l.first = l.first.Prev
		l.first.Next = nil
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
