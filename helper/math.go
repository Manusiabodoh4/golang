package helper

func MathAdd(numbers ...int) int {
	var hasil int
	for _, number := range numbers {
		hasil += number
	}
	return hasil
}
