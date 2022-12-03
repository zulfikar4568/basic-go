package gogoroutines

import (
	"fmt"
	"testing"
	"time"
)

func PrintNama(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Zulfikar Isnaen"
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go PrintNama(channel1)
	go PrintNama(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}
