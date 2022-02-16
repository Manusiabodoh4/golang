package main

import "fmt"

type Blacklist func(string) bool

//function main
func main() {

	sayHello()

	first, last := getFullname()

	first2, last2 := getFullname2()

	sayHelloTo(first, last)
	sayHelloTo(first2, last2)

	fmt.Println(getHello("Reza"))

	total := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println("Total ", total)

	//function value
	goodBye := sayGoodBye
	fmt.Println(goodBye())

	//use function with func paramater
	sayHelloWithFilter("Reza Andriansyah", filterKataKasar)

	//implementation anonymous function
	blacklistFunction := func(name string) bool {
		return name == "Reza"
	}

	loginSystem("Reza", blacklistFunction)

	//another implementation anonymous function
	loginSystem("Amelia", func(name string) bool {
		return name == "Reza"
	})

	fmt.Println("Faktorial 5 :", factorial(5))

}

//function recursive
func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

//function as parameter atau callback with alias
func loginSystem(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Kamu di bloked", name)
		return
	}
	fmt.Println("Selamat datang", name)
}

//function as parameter atau callback
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func filterKataKasar(name string) string {
	if name == "Kasar" {
		return "..."
	}
	return name
}

func sayGoodBye() string {
	return "Selamat tinggal"
}

//function variadic atau functn yang memiliki paramter dinamis (bebas insert berapa aja)
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//function named return value
func getFullname2() (firstname string, lastname string) {
	firstname = "Reza"
	lastname = "Andriansyah"
	return
}

//function return multiple value
func getFullname() (string, string) {
	return "Reza", "Andriansyah"
}

//function return value
func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	}
	return "Hello " + name
}

//function with parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

//function
func sayHello() {
	fmt.Println("Hello semua!!")
}
