package main

import "fmt"

func main() {
	var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range numbers {
		fmt.Println(evenOrOdd(numbers[i]))
	}
}

func evenOrOdd(num int) string {
	if num%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
