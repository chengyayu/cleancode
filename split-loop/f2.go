package main

import (
	"fmt"
	"math"
)

func getYoungestAgeAndTotalSalaryV2(peoples []*People) string {

	youngestAge := getYoungestAge(peoples)

	totalSalary := getTotalSalary(peoples)

	return fmt.Sprintf("youngestAge: %d, totalSalary: %d", youngestAge, totalSalary)
}

func getTotalSalary(peoples []*People) int {
	totalSalary := 0
	for _, p := range peoples {
		totalSalary += p.Salary
	}
	return totalSalary
}

func getYoungestAge(peoples []*People) int {
	youngestAge := math.MaxInt
	for _, p := range peoples {
		if p.Age < youngestAge {
			youngestAge = p.Age
		}
	}
	return youngestAge
}
