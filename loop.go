package main

import "fmt"

func main() {

	//while
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan While Loop ke", counter)
		counter++
	}

	//for
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan For Loop ke", counter)
	}

	//for with range
	slice := []string{"Reza", "Amelia", "Andriansyah", "Rosa"}

	for index, name := range slice {
		fmt.Println("Perulangan For ke ", index, "Dengan value", name)
	}

	dataMap := make(map[string]string)

	dataMap["name"] = "Reza Andriansyah"
	dataMap["alamat"] = "Kota Bekasi"
	dataMap["pekerjaan"] = "Programmer dan Mahasiswa"

	for key, data := range dataMap {
		fmt.Println("Perulangan For key", key, " = ", data)
	}

}
