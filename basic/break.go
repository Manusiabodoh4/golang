package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		if i == 5 {
			break
		}
		i++
	}
	fmt.Println("Index", i)
}
