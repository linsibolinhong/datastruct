package cond

import (
	"sync"
	"testing"
	"time"
)

func TestSingleFlight(t *testing.T) {
	fn := func() (interface{}, error){
		time.Sleep(time.Second*1)
		t.Log("call fn")
		return "wetack", nil
	}

	f := SingleFlight{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond * 100)
		go func() {
			wg.Add(1)
			defer wg.Done()
			t.Log(f.Do("abc", fn))
		}()
	}
	wg.Wait()
}
