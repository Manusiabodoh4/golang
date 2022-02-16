package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func (address Address) PrintElement() {
	//address pada method tersebut hanya lah hasil clone sehingga tidak ref ke pada data di struct nya
	fmt.Println(address.City)
	fmt.Println(address.Province)
	fmt.Println(address.Country)
}

func (address *Address) ChangeToFormatJakarta() {
	address.City = "Jakarata"
	address.Province = "DKI Jakarata"
	address.Country = "Indonesia"
}

func ChangeCountryToAmericaByValue(address Address) {
	//this function just change clone of variable address from paramater
	address.Country = "America"
}

func ChangeCountryToAmericaByRef(address *Address) {
	address.Country = "America"
}

func main() {

	var name string = "Reza Andriansyah"

	fmt.Println(&name)

	var name2 *string = &name

	fmt.Println(name2)

	var address1 Address = Address{"Kota Bekasi", "Jawa Barat", "Indonesia"}
	//melakukan clone atau ref by value
	var address2 Address = address1

	address2.City = "Jakarta"
	address2.Province = "DKI Jakarta"

	//melakukan ref by ref, sehingga ketika melakukan perubahan variable address 1 juga berubah
	var address3 *Address = &address1

	address3.City = "Banten"
	address3.Province = "Jawa Barat"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(*address3)

	//melakukan pembuatan variable bedasarkan pointer dari tipe data yang di tuju
	var address4 *Address = new(Address)

	address4.City = "Tanggerang"
	address4.Province = "Jawa Barat"
	address4.Country = "Indonesia"

	fmt.Println(address4)

	var address5 Address = Address{"Jakarta", "DKI Jakarata", "Indonesia"}

	ChangeCountryToAmericaByValue(address5)

	fmt.Println(address5)

	ChangeCountryToAmericaByRef(&address5)

	fmt.Println(address5)

	var address6 Address = Address{"HUaaaw", "Chicago", "America"}

	address6.PrintElement()

	address6.ChangeToFormatJakarta()

	address6.PrintElement()

}
