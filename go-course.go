package main

import (
	"fmt"
	"time"
)

func push(slice *[]int, elem int) {
	*slice = append(*slice, elem)
}

func pop(slice *[]int) int {
	len := len(*slice)
	r := (*slice)[len-1]
	*slice = (*slice)[:len-1]
	return r
}

func size(slice []int) int {
	return len(slice)
}

func kk(slice []int) {
	slice[0] = 33
	fmt.Printf("IN %v\n", slice)
	slice = append(slice, 44)
	fmt.Printf("IN %v\n", slice)
}

func kk2(slice []int) {
	slice = append(slice, 44)
	fmt.Printf("IN %v\n", slice)
	slice[0] = 33
	fmt.Printf("IN %v\n", slice)
}

func kkborra(slice []int) {
	s := slice[:1]
	fmt.Printf("IN %v\n", s)
}

func f() {
	var prints []func()
	for i := 0; i < 3; i++ {
		prints = append(prints, func() {
			a := i
			fmt.Println(a)
		})
	}
	for _, p := range prints {
		p()
	}
	fmt.Println(time.Now().String())
}

func main() {
	slice := []int{1, 2, 3}
	fmt.Printf("%v\n", slice)
	push(&slice, 42)
	fmt.Printf("%v\n", slice)
	fmt.Printf("%v\n", size(slice))
	r := pop(&slice)
	fmt.Printf("%v %v\n", slice, r)

	fmt.Printf("%v\n", size(slice))

	fmt.Println("--------------------------")

	fmt.Printf("%v\n", slice)
	kk(slice)
	fmt.Printf("%v\n", slice)

	fmt.Println("--------------------------")
	slice = []int{1, 2, 3}
	fmt.Printf("%v\n", slice)
	kk2(slice)
	fmt.Printf("%v\n", slice)

	fmt.Println("--------------------------")
	slice = []int{1, 2, 3}
	fmt.Printf("%v\n", slice)
	kkborra(slice)
	fmt.Printf("%v\n", slice)

	f()
}
