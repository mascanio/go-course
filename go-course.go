package main

import (
	"fmt"
	"strconv"
	"strings"
)

// type Year struct {
// Months [12]Month
// }

// func (year Year) String() string {

// }

type Month struct {
	Weeks []Week
	Name  string
}

func (month Month) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintln(month.Name))
	sb.WriteString(fmt.Sprintln("Mo Tu We Th Fr Sa Su"))
	for _, week := range month.Weeks {
		sb.WriteString(fmt.Sprintf("%v\n", week))
	}

	return sb.String()
}

type Week struct {
	Days [7]Day
}

func (week Week) String() string {
	var sb strings.Builder
	for _, d := range week.Days {
		sb.WriteString(d.String())
		sb.WriteString(" ")
	}
	return sb.String()
}

type Day int

func (day Day) String() string {
	switch {
	case day == 0:
		return "  "
	case day < 0:
		return fmt.Sprint(-day)
	}
	return strconv.Itoa(day)
}

func main() {
	m := Month{
		Name: "July",
		Weeks: []Week{
			{[7]Day{1, 2, 3, 4, 5, 6, 7}},
			{[7]Day{8, 9, 10, 11, 12, 13, 14}},
			{[7]Day{1, 2, 3, 4, 5, 6, 7}},
			{[7]Day{1, 2, 3, 4, 5, 6, 7}},
		},
	}
	fmt.Println(m)
}
