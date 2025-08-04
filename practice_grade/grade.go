package main

import (
	"math/rand"
	"strconv"
)

type score []int
type grade []string
type student []string

func newScore() score {
	scores := score{}
	for i := 0; i < 5; i++ {
		randNumber := rand.Intn(101)
		scores = append(scores, randNumber)
	}

	return scores
}

func (s score) getGrade() grade {
	grades := grade{}
	for i := range s {
		if s[i] >= 80 {
			grades = append(grades, "A")
		} else if s[i] > 69 && s[i] <= 79 {
			grades = append(grades, "B")
		} else if s[i] > 59 && s[i] <= 69 {
			grades = append(grades, "C")
		} else if s[i] > 49 && s[i] <= 59 {
			grades = append(grades, "D")
		} else {
			grades = append(grades, "F")
		}
	}
	return grades
}

func printResult(s score, g grade) student {
	students := student{}
	for i := 1; i <= 5; i++ {
		message := "Student number " + strconv.Itoa(i) + " | score = " + strconv.Itoa(s[i-1]) + " | grade = " + g[i-1]
		students = append(students, message)
	}
	return students
}
