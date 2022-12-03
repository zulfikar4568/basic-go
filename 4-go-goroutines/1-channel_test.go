package gogoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Zulfikar Isnaen"
		// channel <- "Zulfikar Isnaen2" // Tidak bisa mengirim data lagi, default nya hanya 1x, harus ada yang mengkonsumsi dlu
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(2 * time.Second)
}
