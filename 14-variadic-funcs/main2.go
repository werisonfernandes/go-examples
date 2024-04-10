package main

import "fmt"

func variadic(nums ...int) (int, int) {
	var total = 0
	for i, n := range nums {
		fmt.Println(i)
		total += n
	}
	fmt.Println(total)
	return 0, total
}

func main() {
	var a, b = variadic(1, 2, 3, 4, 5, 6)
	fmt.Println(a, b)

	var array = []int{1, 2, 3, 4}
	var _, c = variadic(array...)
	fmt.Println(c)
}
