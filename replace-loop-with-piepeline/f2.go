package main

import (
	"strings"

	"github.com/samber/lo"
)

func acquireDataV2(input string) []CityPhone {
	ls := LineStream(strings.Split(input, "\n"))

	return ls.Slice(1, len(ls)).
		Filter(func(item Line) bool { return !item.IsEmpty() }).
		Map(func(item Line) Item { return item.ToItem() }).
		Filter(func(item Item) bool { return item.Country == "India" }).
		Map(func(item Item) CityPhone { return item.ToCityPhone() })
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

func LineStream(s []string) Lines {
	return lo.Map(s, func(item string, _ int) Line {
		return Line(item)
	})
}

func (p Lines) Slice(start, end int) Lines {
	return lo.Slice(p, start, end)
}

func (p Lines) Filter(f func(item Line) bool) Lines {
	return lo.Filter(p, func(item Line, _ int) bool {
		return f(item)
	})
}

func (p Lines) Map(f func(item Line) Item) List {
	return lo.Map(p, func(item Line, _ int) Item {
		return f(item)
	})
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

func (p List) Filter(f func(item Item) bool) List {
	return lo.Filter(p, func(item Item, _ int) bool {
		return f(item)
	})
}

func (p List) Map(f func(item Item) CityPhone) []CityPhone {
	return lo.Map(p, func(item Item, _ int) CityPhone {
		return f(item)
	})
}
