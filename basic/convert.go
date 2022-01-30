package main

import "fmt"

func main() {

	var nilai int8 = 100
	var nilai64 int64 = int64(nilai)

	var name string = "Reza Andriansyah"
	var karakter uint8 = name[0]

	fmt.Printf("%d = %s\n", karakter, string(karakter))

	fmt.Println(nilai64)

}
