package list

type ListElem struct {
	prev  *ListElem
	next  *ListElem
	list  *List
	Value interface{}
}

type List struct {
	head *ListElem
	tail *ListElem
	len  int
}

func newListElem(v interface{}) *ListElem {
	return &ListElem{
		prev:  nil,
		next:  nil,
		Value: v,
	}
}

func (e *ListElem) Prev() *ListElem {
	return e.prev
}

func (e *ListElem) Next() *ListElem {
	return e.next
}

func New() *List {
	l := List{}
	return l.Init()
}

func (l *List) Init() *List {
	l.head = nil
	l.tail = nil
	l.len = 0
	return l
}

func (l *List) insertElemAfter(e, at *ListElem) *ListElem {
	defer func() {
		l.len++
	}()

	if at == nil {
		l.head = e
		l.tail = e
		e.list = l
		return e
	}

	e.list = l
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	if n != nil {
		n.prev = e
	} else {
		l.tail = e
	}

	return e
}

func (l *List) insertElemBefore(e, at *ListElem) *ListElem {
	defer func() {
		l.len++
	}()

	if at == nil {
		l.head = e
		l.tail = e
		e.list = l
		return e
	}

	e.list = l
	n := at.prev
	at.prev = e
	e.next = at
	e.prev = n
	if n != nil {
		n.next = e
	} else {
		l.head = e
	}

	return e
}

func (l *List) Remove(e *ListElem) *ListElem {
	if e == nil || e.list != l {
		return e
	}

	if e.prev != nil {
		e.prev.next = e.next
	} else {
		l.head = e.next
	}

	if e.next != nil {
		e.next.prev = e.prev
	} else {
		l.tail = e.prev
	}

	e.list = nil
	l.len--
	return e
}

func (l *List) PushFront(v interface{}) *ListElem {
	return l.insertElemBefore(newListElem(v), l.head)
}

func (l *List) PushBack(v interface{}) *ListElem {
	return l.insertElemAfter(newListElem(v), l.tail)
}

func (l *List) InsertBefore(v interface{}, mark *ListElem) *ListElem {
	if mark == nil || mark.list != l {
		return nil
	}
	return l.insertElemBefore(newListElem(v), mark)
}

func (l *List) InsertAfter(v interface{}, mark *ListElem) *ListElem {
	if mark == nil || mark.list != nil {
		return nil
	}
	return l.insertElemAfter(newListElem(v), mark)
}

func (l *List) Swap(a, b *ListElem) {
	if a == nil || a.list != l || b == nil || b.list != l {
		return
	}

	if a == b {
		return
	}

	if b.next == a {
		a, b = b, a
	}

	pa, pb := a.prev, b.prev
	na, nb := a.next, b.next

	if a.next == b {
		a.next = b.next
		if b.next != nil {
			b.next.prev = a
		}
		if a.prev != nil {
			a.prev.next = b
		}
		a.prev = b
		b.next = a
		b.prev = pa
	} else {
		if pa != nil {
			pa.next = b
		}

		if pb != nil {
			pb.next = a
		}

		if na != nil {
			na.prev = b
		}

		if nb != nil {
			nb.prev = a
		}

		a.prev, b.prev = pb, pa
		a.next, b.next = nb, na
	}

	if a.next == nil {
		l.tail = a
	}

	if b.next == nil {
		l.tail = b
	}

	if a.prev == nil {
		l.head = a
	}

	if b.prev == nil {
		l.head = b
	}
}

func (l *List) Front() *ListElem {
	return l.head
}

func (l *List) Back() *ListElem {
	return l.tail
}

func (l *List) Len() int {
	return l.len
}

func (l *List) At(idx int) *ListElem {
	if idx >= l.Len() {
		return nil
	}

	item := l.Front()
	for i := 0; i < idx; i++ {
		item = item.Next()
	}
	return item
}

func (l *List) BackAt(idx int) *ListElem {
	if idx >= l.Len() {
		return nil
	}

	item := l.Back()
	for i:=0; i < idx; i++ {
		item = item.Prev()
	}
	return item
}