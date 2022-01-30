package main

import "fmt"

func main() {

	name1 := "Reza Andriansyah"
	name2 := "Amelia Rosa"

	res := name1 == name2

	fmt.Println(res)

	nilai1 := 100
	nilai2 := 90

	fmt.Println(nilai1 > nilai2)
	fmt.Println(nilai1 < nilai2)
	fmt.Println(nilai1 >= nilai2)
	fmt.Println(nilai1 >= nilai2)
	fmt.Println(nilai1 != nilai2)

	kondisi1 := false
	kondisi2 := true

	fmt.Println(kondisi1 && kondisi2)
	fmt.Println(kondisi1 || kondisi2)
	fmt.Println(!kondisi1)
	fmt.Println(!kondisi2)

}
