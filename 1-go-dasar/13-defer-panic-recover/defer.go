package main

import "fmt"

func main() {
	aplikasiRunning(0)
}

func logging() {
	fmt.Println("Lakukan Logging")
}

func aplikasiRunning(value int) {
	// defer digunakan untuk menjalankan fungsi ketika fungsi aplikasiRunning selesai dijalankan
	// defer akan tetep di eksekusi walaupun program error
	defer logging()

	fmt.Println("Aplikasi berjalan")
	errorNih := 2 / value
	fmt.Println("Result :", errorNih)

}
