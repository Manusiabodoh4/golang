package main

import "fmt"

func main() {

	var nilai int8 = 88

	var nilai2 int8 = 98

	if nilai > nilai2 {
		fmt.Println("Nilai lebih besar dari pada nilai2")
	} else {
		fmt.Println("Nilai2 lebih besar dari pada nilai")
	}

	nama := "Reza"

	if nama == "Reza" {
		fmt.Println("Hello Reza")
	} else if nama == "Amelia" {
		fmt.Println("Halo Sayang ku!")
	} else {
		fmt.Println("Hi, boleh kenala?")
	}

	panjangNama := len(nama)

	if panjangNama > 5 {
		fmt.Println("Panjang sekali yaa")
	}

	//shortpass
	if lengNama := len(nama); lengNama > 5 {
		fmt.Println("Panjang sekali yaa")
	}

}
