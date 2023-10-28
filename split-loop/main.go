package main

import (
	"fmt"
)

func main() {
	peoples := []*People{
		{Age: 22, Salary: 100},
		{Age: 30, Salary: 80},
		{Age: 35, Salary: 90},
	}
	rest1 := getYoungestAgeAndTotalSalaryV1(peoples)
	rest2 := getYoungestAgeAndTotalSalaryV2(peoples)
	fmt.Println(rest1, rest2)
}

type People struct {
	Age    int
	Salary int
}
