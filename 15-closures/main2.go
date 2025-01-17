package main

import "fmt"

func intSeq() func() int {
	var i = 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	var nextInt = intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
