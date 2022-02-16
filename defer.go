package main

import "fmt"

//defer merupakan sebuah fungsi yang dijadwalkan akan dijalankan ketika fungsi tersebut telah selesai dijalankan

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	//implmenetation defer
	defer logging()
	fmt.Println("Run Application")
}

func runApplicationMath(num int) {
	defer logging()
	fmt.Println("Run Application Math")
	hasil := 10 / num
	fmt.Println("Hasil = ", hasil)
}

func main() {
	runApplication()
	runApplicationMath(10)
	//walaupun menghasilkan error defer akan tetap berjalan
	runApplicationMath(0)
}
