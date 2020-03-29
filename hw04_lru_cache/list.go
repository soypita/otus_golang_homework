package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *listItem                  // первый Item
	Back() *listItem                   // последний Item
	PushFront(v interface{}) *listItem // добавить значение в начало
	PushBack(v interface{}) *listItem  // добавить значение в конец
	Remove(i *listItem)                // удалить элемент
	MoveToFront(i *listItem)           // переместить элемент в начало
}

type listItem struct {
	Value interface{}
	Next  *listItem
	Prev  *listItem
}

type list struct {
	FrontItem *listItem
	BackItem  *listItem
	length    int
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *listItem {
	return l.FrontItem
}

func (l *list) Back() *listItem {
	return l.BackItem
}

func (l *list) insertBefore(node *listItem, newNode *listItem) {
	newNode.Prev = node
	if node.Next == nil {
		newNode.Next = nil
		l.FrontItem = newNode
	} else {
		newNode.Next = node.Next
		node.Next.Prev = newNode
	}
	node.Next = newNode
}

func (l *list) insertAfter(node *listItem, newNode *listItem) {
	newNode.Next = node
	if node.Prev == nil {
		newNode.Prev = nil
		l.BackItem = newNode
	} else {
		newNode.Prev = node.Prev
		node.Prev.Next = newNode
	}
	node.Prev = newNode
}

func (l *list) PushFront(v interface{}) *listItem {
	item := &listItem{}
	item.Value = v
	if l.FrontItem == nil {
		l.FrontItem = item
		l.BackItem = item
	} else {
		l.insertBefore(l.FrontItem, item)
	}

	l.length++
	return item
}

func (l *list) PushBack(v interface{}) *listItem {
	if l.BackItem == nil {
		return l.PushFront(v)
	}
	item := &listItem{}
	item.Value = v

	l.insertAfter(l.BackItem, item)
	l.length++
	return item
}

func (l *list) Remove(i *listItem) {
	if i == nil {
		return
	}
	prevItem := i.Prev
	nextItem := i.Next
	if nextItem != nil {
		nextItem.Prev = prevItem
	} else {
		l.FrontItem = prevItem
	}

	if prevItem != nil {
		prevItem.Next = nextItem
	} else {
		l.BackItem = nextItem
	}
	l.length--
}

func (l *list) MoveToFront(i *listItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return &list{}
}
