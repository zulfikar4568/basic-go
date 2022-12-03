package gogoroutines

import (
	"fmt"
	"testing"
)

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "M"
	channel <- "Zulfikar"
	channel <- "Isnaen"

	fmt.Println(cap(channel))
	fmt.Println(len(channel))

	fmt.Println(<-channel)
	fmt.Println(len(channel))
	fmt.Println(<-channel)
	fmt.Println(len(channel))
	fmt.Println(<-channel)
	fmt.Println(len(channel))

	fmt.Println("Finish")
}
