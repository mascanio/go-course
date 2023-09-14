package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
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

const (
	as = iota + 1
	dos
	tres
	cuatro
	cinco
	seis
	siete
	ocho
	nueve
	diez = 10
	jack
	reina
	rey
)

func blackjack(a int, b int) int {
	if a+b+10 <= 21 {
		// Promote as if present
		if a == as {
			return b + 11
		}
		if b == as {
			return a + 11
		}
	}
	return a + b
}

func multiplos(lim int) int {
	acu := 0
	for n := 1; n*3 < lim; n++ {
		acu += n * 3
		if n*5 < lim {
			acu += n * 5
		}
	}
	return acu
}

func multiplos_2(lim int) int {
	acu := 0
	for i := 0; i < lim; i++ {
		if i%3 == 0 {
			acu += i
		}
		if i%5 == 0 {
			acu += i
		}
	}
	return acu
}

func filter_str_rune_bigger_1_byte_2(s string) string {
	rune_mapper := func(r rune) rune {
		accept_rune := rune_filter(r)
		if accept_rune {
			return r
		}
		return -1
	}
	return strings.Map(rune_mapper, s)
}
func filter_str_rune_bigger_1_byte(s string) string {
	var r string
	for _, rune := range s {
		if rune_filter(rune) {
			r = r + string(rune)
		}
	}

	return r
}

func rune_filter(r rune) bool {
	return utf8.RuneLen(r) > 1
}

func main() {
	// 1
	celsius := 37.3
	fmt.Printf("%v grados centígrados son %v grados Farenheit\n", celsius, celsius_to_fahrenheit(float64(celsius)))
	// 2
	fmt.Printf("%v\n", is_square_triangle(5, 3, 4))
	// 3
	fmt.Printf("%v\n", position(-1, -11))
	// 4
	fmt.Printf("El valor final es %v\n", blackjack(as, rey))
	fmt.Printf("El valor final es %v\n", blackjack(ocho, rey))
	fmt.Printf("El valor final es %v\n", blackjack(as, as))
	// 5
	fmt.Printf("La suma es %v\n", multiplos(1000000))
	fmt.Printf("La suma es %v\n", multiplos_2(1000000))
	// 6
	str_1 := "Hola 🗺 Que seas muy 🙋 para siempre"
	fmt.Printf("In: %v\nOut: %v\n", str_1, filter_str_rune_bigger_1_byte(str_1))
	fmt.Printf("In: %v\nOut: %v\n", str_1, filter_str_rune_bigger_1_byte_2(str_1))
}
