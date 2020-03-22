package list

import (
	"fmt"
	"testing"
)

func testNil(l *List, t *testing.T) {
	if l.Front() != nil || l.Back() != nil || l.Len() != 0 {
		t.Errorf("Not nil")
	}
}

func testSingle(l *List, val interface{}, t *testing.T) {
	front := l.Front()
	if front.Value != val || front.Prev() != nil || front.Next() != nil {
		t.Errorf("front fail")
	}

	back := l.Back()
	if back.Value != val || back.Prev() != nil || back.Next() != nil {
		t.Errorf("back fail")
	}
}

func testMult(l *List, vals []interface{}, t *testing.T) {
	if l.Len() != len(vals) {
		t.Errorf("Len mismatch")
		return
	}

	front := l.Front()
	for idx := 0; idx < len(vals); idx++ {
		if front.Value != vals[idx] {
			t.Errorf("Mis Val %+v:%+v", front.Value, vals[idx])
			return
		}
		//t.Log(front.Value)
		front = front.Next()
	}
	if front != nil {
		t.Errorf("Circle front")
	}

	back := l.Back()
	for idx := len(vals) - 1; idx >= 0; idx-- {
		if back.Value != vals[idx] {
			t.Errorf("Mis Val %+v:%+v", front.Value, vals[idx])
			return
		}
		back = back.Prev()
	}

	if back != nil {
		t.Errorf("Circle Back")
	}
}


func TestList(t *testing.T) {
	l := New()
	testNil(l, t)

	l.PushFront("head")
	testSingle(l, "head", t)

	l.Remove(l.Front())
	testNil(l, t)

	l.PushBack("back")
	testSingle(l, "back", t)

	l.Remove(l.Back())
	testNil(l, t)

	l.Init()
	k := 10
	a := []interface{}{}
	for i := 0; i < k; i++ {
		a = append(a, i)
		l.PushBack(i)
	}
	testMult(l, a, t)

	l.Init()
	a = []interface{}{}
	for i := 0; i < k; i++ {
		a = append([]interface{}{i}, a...)
		l.PushFront(i)
	}
	testMult(l, a, t)


	l.Init()
	a = []interface{}{}
	for i := 0; i < k; i++ {
		a = append(a, i)
		l.PushBack(i)
	}
	l.Swap(l.Front(), l.Back())
	l.Swap(l.At(5), l.At(3))
	l.Swap(l.At(1), l.At(2))
	a[0] = k-1
	a[k-1] = 0

	a[1], a[2] = a[2], a[1]
	a[3], a[5] = a[5], a[3]

	testMult(l, a, t)
	cnt := 1
	for l.Len() != 0 || l.Front() != nil {
		fmt.Println(l.Front().Value, a[k - l.Len()], l.Len())
		l.Remove(l.Front())
		cnt ++
		if cnt > 20 {
			return
		}
	}
}
