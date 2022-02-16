package main

import "fmt"

func main() {

	//multiple deklarasi konstanta
	const (
		nilai int8 = 100
		kelas int8 = 12
	)

	//multiple deklarasi variable
	var (
		firstName string
		lastName  string
	)

	//deklarasi variable
	var name string

	//konstanta
	const age uint8 = 0

	//konstanta tanpa deklarasi tipe
	const gender = "Laki Laki"

	//membuat tanpa menggunakan keyword var
	alamat := "Kota Bekasi"

	name = "Reza Andriansyah"
	alamat = "Kota Jakarta"

	firstName = "Reza"
	lastName = "Andriansyah"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(gender)
	fmt.Println(alamat)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
