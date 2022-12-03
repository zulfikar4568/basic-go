package gogoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsync(group *sync.WaitGroup) {
	defer group.Done() // Jika kita tidak menambahkan ini maka akan terjadi deadlock

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsync(group)
	}

	group.Wait()
	fmt.Println("Complete")
}
