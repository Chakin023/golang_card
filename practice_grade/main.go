package main

import "fmt"

func main() {
	scores := newScore()
	fmt.Println(scores)
	grades := scores.getGrade()
	fmt.Println(grades)
	results := printResult(scores, grades)
	for i := range results {
		fmt.Println(results[i])
	}
}
