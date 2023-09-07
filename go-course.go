package main

import (
	"fmt"
	"math"
)

func celsius_to_fahrenheit(celsius float64) float64 {
	return (celsius * 9.0 / 5.0) + 32.0
}

func catetos_hipotenusa(d1 int, d2 int, d3 int) (int, int, int) {
	switch {
	case d1 > d2 && d1 > d3:
		return d3, d2, d1
	case d2 > d1 && d2 > d3:
		return d1, d3, d2
	}
	return d1, d2, d3
}

func is_square_triangle(d1 int, d2 int, d3 int) bool {
	c1, c2, h := catetos_hipotenusa(d1, d2, d3)
	return math.Pow(float64(c1), 2)+math.Pow(float64(c2), 2) == math.Pow(float64(h), 2)
}

func position(lat float64, long float64) string {
	switch {
	case lat == 0.0:
		switch {
		case long == 0.0:
			return "C"
		case long < 0.0:
			return "W"
		case long > 0.0:
			return "E"
		}
	case lat < 0.0:
		switch {
		case long == 0.0:
			return "S"
		case long < 0.0:
			return "SW"
		case long > 0.0:
			return "SE"
		}
	case lat > 0.0:
		switch {
		case long == 0.0:
			return "N"
		case long < 0.0:
			return "NW"
		case long > 0.0:
			return "NE"
		}
	}
	return "error"
}

func main() {
	// 1
	celsius := 37.3
	fmt.Printf("%v grados centígrados son %v grados Farenheit\n", celsius, celsius_to_fahrenheit(float64(celsius)))
	// 2
	fmt.Printf("%v\n", is_square_triangle(5, 3, 4))
	// 3
	fmt.Printf("%v\n", position(-1, -11))
}
