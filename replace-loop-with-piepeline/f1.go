package main

import (
	"strings"
)

func acquireDataV1(input string) []CityPhone {
	lines := strings.Split(input, "\n")
	firstLine := true
	var result []CityPhone
	for _, line := range lines {
		if firstLine {
			firstLine = false
			continue
		}
		if strings.TrimSpace(line) == "" {
			continue
		}
		record := strings.Split(line, ",")
		if strings.TrimSpace(record[1]) == "India" {
			result = append(result, CityPhone{
				City:  strings.TrimSpace(record[0]),
				Phone: strings.TrimSpace(record[2]),
			})
		}
	}
	return result
}
