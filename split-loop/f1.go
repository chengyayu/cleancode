package main

import (
	"fmt"
	"math"
)

func getYoungestAgeAndTotalSalaryV1(peoples []*People) string {
	var youngestAge int
	if peoples[0] != nil {
		youngestAge = peoples[0].Age
	} else {
		youngestAge = math.MaxInt
	}

	var totalSalary = 0
	for _, p := range peoples {
		if p.Age < youngestAge {
			youngestAge = p.Age
		}
		totalSalary += p.Salary
	}

	return fmt.Sprintf("youngestAge: %d, totalSalary: %d", youngestAge, totalSalary)
}
