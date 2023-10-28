package main

import (
	"fmt"
	"math"
)

func getYoungestAgeAndTotalSalaryV2(peoples []*People) string {

	youngestAge := math.MaxInt
	for _, p := range peoples {
		if p.Age < youngestAge {
			youngestAge = p.Age
		}
	}

	totalSalary := 0
	for _, p := range peoples {
		totalSalary += p.Salary
	}

	return fmt.Sprintf("youngestAge: %d, totalSalary: %d", youngestAge, totalSalary)
}
