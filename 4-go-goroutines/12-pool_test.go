package gogoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var pool = sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	var s1, s2, s3 = "M", "Zulfikar", "Isnaen"

	pool.Put(&s1)
	pool.Put(&s2)
	pool.Put(&s3)

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
}
