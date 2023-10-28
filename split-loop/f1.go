package main

import (
	"fmt"
	"math"
)

func getYoungestAgeAndTotalSalaryV1(peoples []*People) string {
	var yougestAge int
	if peoples[0] != nil {
		yougestAge = peoples[0].Age
	} else {
		yougestAge = math.MaxInt
	}

	var totalSalary = 0
	for _, p := range peoples {
		if p.Age < yougestAge {
			yougestAge = p.Age
		}
		totalSalary += p.Salary
	}

	return fmt.Sprintf("youngestAge: %d, totalSalary: %d", yougestAge, totalSalary)
}
