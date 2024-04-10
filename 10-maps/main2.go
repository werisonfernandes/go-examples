package main

import "fmt"

func main() {
	var m = make(map[string]int)
	m["k1"] = 2
	m["k2"] = 12
	fmt.Println(m)

	delete(m, "k2")
	fmt.Println(m)

	_, prs := m["k1"]
	fmt.Println("prs:", prs)

	var mm = map[string]int{}
	fmt.Println("prs:", mm)
}
