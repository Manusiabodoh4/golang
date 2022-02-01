package main

import "fmt"

func main() {

	counter := 0

	increment := func() {
		fmt.Println("Increment")
		//implementation closure = variable dapat dirubah didalam sebuah function, namun variable yang ada di function tidak dapat diakses di luar function
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)

}
