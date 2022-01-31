package main

import "fmt"

func main() {

	var nilai = [...]int8{80, 90, 99, 75, 100}

	var slice1 = nilai[2:4]
	var slice2 = nilai[2:]
	var slice3 = nilai[:4]
	var slice4 = nilai[:]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	fmt.Printf("Panjang Slice1 : %d\n", len(slice1))
	fmt.Printf("Kapasitas Slice 1 : %d\n", cap(slice1))

	//ketika slice diruba, data dalam array juga berubah (sifat change by address/pointer)
	slice1[0] = 98

	fmt.Println(nilai)

	sliceAppend := append(slice1, 77)

	fmt.Println(sliceAppend)

	//membuat slice menggunakan method make()
	sliceMake := make([]string, 2, 5)

	sliceMake[0] = "Reza"
	sliceMake[1] = "Andriansyah"

	fmt.Println(sliceMake)
	fmt.Println(len(sliceMake))
	fmt.Println(cap(sliceMake))

	sliceMake = append(sliceMake, "Amelia")

	fmt.Println(sliceMake)

	sliceCopy := make([]string, len(sliceMake), cap(sliceMake))

	copy(sliceCopy, sliceMake)

	fmt.Println(sliceCopy)

	//array vs slice
	iniArray := [...]int8{1, 2, 3, 4, 5}
	iniSlice := []int8{1, 2, 3, 4, 5}

	//array
	fmt.Println(iniArray)

	//slice
	fmt.Println(iniSlice)

}
