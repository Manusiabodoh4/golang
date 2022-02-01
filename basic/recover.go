package main

import "fmt"

func endApp() {
	fmt.Println("Applikasi selesai dijalankan")
	message := recover()
	fmt.Println("Terjadi error dengan message :", message)
}

func runApp(err bool) {
	defer endApp()
	if err {
		panic("Testing Error")
	}
}

func main() {
	runApp(true)
}
