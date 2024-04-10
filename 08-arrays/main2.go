package main

import "fmt"

func main() {
	var a []int
	fmt.Println(len(a))
	fmt.Println(a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var c = []int{1, 2, 3, 4, 5}
	fmt.Println(c)

	var m = [2][3]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(m)
}
