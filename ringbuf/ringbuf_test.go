package ringbuf

import (
	"fmt"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	r := New(3)
	fmt.Println(r.Empty())
	n := r.Write([]byte("123"))
	fmt.Println(n)
	fmt.Println(r.Detail())
	rd, _ := r.Read(1)
	fmt.Println(string(rd))
	fmt.Println(r.Detail())

	r.Write([]byte("1"))
	fmt.Println(r.Detail())

	rd, _ = r.Read(2)
	fmt.Println(string(rd))
	fmt.Println(r.Detail())

	r.Write([]byte("23"))
	fmt.Println(r.Detail())

	rd, _ = r.ReadWithOffset(1,1)
	fmt.Println(string(rd))
	r.Read(1)
	rd, _ = r.ReadWithOffset(1,1)
	fmt.Println(string(rd))
	fmt.Println(r.Detail())


	r.Clear()
	r.begin = 1
	r.end = 1
	r.Write([]byte("12"))
	fmt.Println(r.Detail())

	r.Remove(1,1)
	fmt.Println(r.Detail())

	r.Remove(0,1)
	fmt.Println(r.Detail())



}
