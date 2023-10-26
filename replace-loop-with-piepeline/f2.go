package main

import (
	"strings"

	"github.com/samber/lo"
)

func acquireDataV2(input string) []CityPhone {
	lines := Lines{}.From(strings.Split(input, "\n"))
	rest := lines.Slice(1, len(lines)).
		Filter(func(item Line, _ int) bool { return !item.IsEmpty() }).
		Map(func(item Line, _ int) Item { return item.ToItem() }).
		Filter(func(item Item, _ int) bool { return item.Country == "India" }).
		Map(func(item Item, _ int) CityPhone { return item.ToCityPhone() })
	return rest
}

type Line string

func (p Line) IsEmpty() bool {
	return string(p) == ""
}

func (p Line) ToItem() Item {
	l := strings.Split(string(p), ",")
	return Item{
		City:    strings.TrimSpace(l[0]),
		Country: strings.TrimSpace(l[1]),
		Phone:   strings.TrimSpace(l[2]),
	}
}

type Lines []Line

func (p Lines) From(s []string) Lines {
	return lo.Map(s, func(item string, _ int) Line {
		return Line(item)
	})
}

func (p Lines) Slice(start, end int) Lines {
	return lo.Slice(p, start, end)
}

func (p Lines) Filter(f func(item Line, _ int) bool) Lines {
	return lo.Filter(p, f)
}

func (p Lines) Map(f func(item Line, _ int) Item) List {
	return lo.Map(p, f)
}

type Item struct {
	Country string
	City    string
	Phone   string
}

func (p Item) ToCityPhone() CityPhone {
	return CityPhone{
		City:  p.City,
		Phone: p.Phone,
	}
}

type List []Item

func (p List) Filter(f func(item Item, _ int) bool) List {
	return lo.Filter(p, f)
}

func (p List) Map(f func(item Item, _ int) CityPhone) []CityPhone {
	return lo.Map(p, f)
}
