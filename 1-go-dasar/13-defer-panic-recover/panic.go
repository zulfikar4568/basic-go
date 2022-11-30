// Panic digunkan untuk menghentikan program
// Panic biasanya digunakan saat program error
// Saat Panic dipanggil, program akan terhenti, tetapi fungsi defer akan tetap di eksekusi

package main

import "fmt"

func main() {
	aplikasiRunning(true)
}

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func aplikasiRunning(error bool) {
	// defer digunakan untuk menjalankan fungsi ketika fungsi aplikasiRunning selesai dijalankan
	// defer akan tetep di eksekusi walaupun program error
	defer endApp()
	if error {
		panic("Aplikasi terjadi Error")
	}
	fmt.Println("Aplikasi Berjalan")
}
