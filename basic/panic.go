package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi telah selesai")
}

func runApp(err bool) {
	defer endApp()
	if err {
		panic("Error")
	}
}

func main() {
	runApp(false)
	runApp(true)
}
