package main

import (
	"fmt"
)

const offices = `office, country, telephone
Chicago, USA, +1 312 373 1000
Beijing, China, +86 4008 900 505
Bangalore, India, +91 80 4064 9570
Porto Alegre, Brazil, +55 51 3079 3550
Chennai, India, +91 44 660 44766`

func main() {
	rest1 := acquireDataV1(offices)
	rest2 := acquireDataV2(offices)
	fmt.Println(rest1, rest2)
}

type CityPhone struct {
	City  string
	Phone string
}
