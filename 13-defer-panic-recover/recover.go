// recover digunakan untuk menangkap data panic
// jika kita menggunakan recover maka kita bisa men skip throw panic, sehingga bisa melanjutkan program selanjutnya

package main

import "fmt"

func main() {
	aplikasiRunning(true)
	fmt.Println("Zulfikar")
}

func endApp() {
	msg := recover()
	if msg != nil {
		fmt.Println("Error", msg)
	}
	fmt.Println("Aplikasi selesai")
}

func aplikasiRunning(error bool) {
	// defer digunakan untuk menjalankan fungsi ketika fungsi aplikasiRunning selesai dijalankan
	// defer akan tetep di eksekusi walaupun program error
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
}
