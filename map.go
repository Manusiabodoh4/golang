package main

import "fmt"

func main() {

	//map digunakan degan menentukan tipe data dari key dan valuenya.
	person := map[string]string{
		"name":   "Reza Andriansyah",
		"alamat": "Kota Bekasi",
	}

	person["pekerjaan"] = "Programmer, Mahasiswa"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["alamat"])

	//untuk mendapatkan jumlah data pada person
	lenMap := len(person)

	fmt.Println(lenMap)

	//membuat map dengan menggunakan method make()
	mapMake := make(map[string]string)

	mapMake["data"] = "ini data"

	fmt.Println(mapMake)

	delete(person, "pekerjaan")

	fmt.Println(person)

}
