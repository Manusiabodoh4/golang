package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello my name is", customer.Name)
}

func main() {

	var customer1 Customer

	customer1.Name = "Reza Andriansyah"
	customer1.Address = "Kota Bekasi"
	customer1.Age = 20

	fmt.Println(customer1)

	cus2 := Customer{
		Name:    "Amelia Rosa",
		Address: "Kota Bekasi",
		Age:     21,
	}

	fmt.Println(cus2)

	cus2.sayHello()

}
