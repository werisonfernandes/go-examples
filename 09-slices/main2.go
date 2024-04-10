package main

import "fmt"

func main() {
	var s = make([]int, 3)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	s = append(s, 3)
	fmt.Println(s)

	var ss = make([][]int, 2, 2)
	ss[0] = make([]int, 2)
	ss[0][0] = 998
	ss[0][1] = 999
	//
	ss[1] = make([]int, 2)
	ss[1][0] = 222
	fmt.Println(ss)
}
