package main

import (
	"fmt"
	"math/rand"
)

func main() {
	realNum := rand.Intn(3) + 1
	var guessNum int
	fmt.Println("Guess the integer from 1 to 3: ")
	_, err := fmt.Scanf("%d", &guessNum)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(guessNumber(guessNum, realNum))
}

func guessNumber(guessNum int, realNum int) string {
	if realNum == guessNum {
		return "Correct!"
	} else {
		return fmt.Sprintf("Wrong! The realNum is %d but your guessNum is %d", realNum, guessNum)
	}
}
