package main

func addMaxMin(max int, min int) int {
	return max + min
}

func addSomeNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
