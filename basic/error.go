package main

import (
	"errors"
	"fmt"
)

//tipe error universal pada golang
type error interface {
	Error() string
}

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi nya adalah nol")
	}
	return (nilai / pembagi), nil
}

func main() {
	hasil, error := Pembagian(10, 0)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	fmt.Println(hasil)
}
