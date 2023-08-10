package main

import (
	"fmt"
)

func main() {
	x := soma(1, 2, 3)
	y := multiplica(10, 10)
	z := divide(15, 15)
	w := subtrai(20, 3)
	fmt.Println(x, y, z, w)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}

func divide(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	total := i[0]
	for _, v := range i[1:] {
		total /= v
	}
	return total
}
func subtrai(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	total := i[0]
	for _, v := range i[1:] {
		total -= v
	}
	return total
}
