package gogoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
	close(channel)
}

// Hanya bisa Kirim data
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Zulfikar Isnaen"
}

// Hanya bisa Menerima data
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}
