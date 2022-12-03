package gogoroutines

import (
	"fmt"
	"testing"
	"time"
)

func KirimData(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Zulfikar Isnaen"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go KirimData(channel)

	data := <-channel
	fmt.Println(data)
}
