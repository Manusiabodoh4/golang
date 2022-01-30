package main

import "fmt"

func main() {

	const (
		nilai1 int8 = 40
		nilai2 int8 = 60
	)

	hasil := nilai1 + nilai2

	hasil += 10

	hasil++

	fmt.Println(hasil)

}
