package main

import "fmt"

func MakeMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}
}

func main() {
	person := MakeMap("Reza Andriansyah")
	if person == nil {
		fmt.Println("Data Kosong")
		return
	}
	fmt.Println(person["name"])
}
