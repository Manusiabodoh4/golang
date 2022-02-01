package main

import "fmt"

func main() {

	name := "Reza"

	switch name {
	case "Amelia":
		fmt.Println("Hello Amelia Rosa")
	case "Reza":
		fmt.Println("Halo Reza Andriansyah")
	default:
		fmt.Println("Halo kamu siapa??")
	}

	umur := 20

	switch {
	case umur >= 17:
		fmt.Println("Sudah Dewasa")
	case umur <= 17:
		fmt.Println("Belum Dewasa")
	}

}
