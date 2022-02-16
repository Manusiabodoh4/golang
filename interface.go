package main

import "fmt"

type InterfacePerson interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return "Nama = " + person.Name
}
func (animal Animal) GetName() string {
	return "Nama Hewan = " + animal.Name
}

func SayHello(data InterfacePerson) {
	fmt.Println(data.GetName())
}

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return false
	}
}

func main() {
	person1 := Person{
		Name: "Reza Andriansyah",
	}
	hewan1 := Animal{
		Name: "Harimau",
	}
	SayHello(person1)
	SayHello(hewan1)
	var data interface{} = Ups(1)
	fmt.Println(data)
}
