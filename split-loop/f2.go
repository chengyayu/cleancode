package main

import (
	"fmt"
	"math"

	"github.com/samber/lo"
)

func getYoungestAgeAndTotalSalaryV2(peoples []*People) string {
	return fmt.Sprintf("youngestAge: %d, totalSalary: %d", getYoungestAge(peoples), getTotalSalary(peoples))
}

func getTotalSalary(peoples []*People) int {
	salarys := lo.Map(peoples, func(item *People, _ int) int {
		return item.Salary
	})
	return lo.Sum(salarys)
}

func getYoungestAge(peoples []*People) int {
	ages := lo.Map(peoples, func(item *People, _ int) int {
		return item.Age
	})
	return lo.Min(append(ages, math.MaxInt))
}
