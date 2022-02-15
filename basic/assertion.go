package main

import "fmt"

//assertion = merubah type pada sebuah variable

func Random() interface{} {
	return "OK"
}

func main() {
	result := Random()
	resultString := result.(string)
	fmt.Println(resultString)
}
