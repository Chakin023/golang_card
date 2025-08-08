package main

import (
	"fmt"
)

type student struct {
	firstName string
	lastName  string
	score     float64
}

func main() {
	john := student{
		firstName: "John",
		lastName:  "Doe",
		score:     50.5,
	}

	jane := student{
		firstName: "Jane",
		lastName:  "Smith",
		score:     60.7,
	}

	mark := student{
		firstName: "Mark",
		lastName:  "Lee",
		score:     90.9,
	}

	allStudent := []*student{&john, &jane, &mark}

	fmt.Println("--- Before ---")
	for _, student := range allStudent {
		student.printResult()
	}

	//can add loop for update each of student score
	allStudent[0].updateScore(55.5)
	allStudent[1].updateScore(66.7)
	allStudent[2].updateScore(99.9)

	fmt.Println("--- After ---")
	for _, student := range allStudent {
		student.printResult()
	}

	fmt.Println("--- Average Score: ---")
	fmt.Println(averageScore(allStudent))

	fmt.Println("--- Max Score: ---")
	maxS := maxScore(allStudent)
	fmt.Printf("For max is %s with score %.2f\n", maxS.firstName, maxS.score)

	fmt.Println("--- Min Score: ---")
	minS := minScore(allStudent)
	fmt.Printf("For min is %s with score %.2f\n", minS.firstName, minS.score)
}

func (s *student) updateScore(newScore float64) {
	s.score = newScore
}

// can change to input the [] and loop inside func to print
func (s student) printResult() {
	fmt.Printf("Name: %s %s | Score: %.2f\n", s.firstName, s.lastName, s.score)
}

func averageScore(studentList []*student) float64 {
	var sum float64
	for _, s := range studentList {
		sum += s.score
	}
	average := float64(sum) / float64(len(studentList))
	return average
}

func maxScore(studentList []*student) *student {
	maxStu := studentList[0]
	for _, s := range studentList {
		if s.score > maxStu.score {
			maxStu = s
		}
	}
	return maxStu
}

func minScore(studentList []*student) *student {
	minStu := studentList[0]
	for _, s := range studentList {
		if s.score < minStu.score {
			minStu = s
		}
	}
	return minStu
}
