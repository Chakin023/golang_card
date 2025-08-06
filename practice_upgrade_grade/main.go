package main

import "fmt"

func main() {
	scores := score{}
	for i := 1; i <= 5; i++ {
		var inputNum int
		fmt.Printf("Enter score for student %d: ", i)
		_, err := fmt.Scanln(&inputNum)
		if err != nil {
			fmt.Println(err)
			i--
			continue
		}
		scores = append(scores, inputNum)
	}

	fmt.Println("--- Results ---")
	grades := scores.getGrade()
	results := printResult(scores, grades)
	for i := range results {
		fmt.Println(results[i])
	}
	fmt.Println("--- Average Score ---")
	fmt.Printf("The average score is %.2f", scores.averageScore())
}
