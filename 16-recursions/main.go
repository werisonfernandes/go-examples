// Go suporta
// <a href="http://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>funções recursivas</em></a>.
// Aqui temos um clássico exemplo de fatorial.

package main

import "fmt"

// Essa função `fact` chama a si mesma até atingir o
// caso base de `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(0))
	fmt.Println(fact(1))
	fmt.Println(fact(2))
	fmt.Println(fact(3))
	fmt.Println(fact(4))
	fmt.Println(fact(5))
	fmt.Println(fact(6))
	fmt.Println(fact(7))
	fmt.Println(fact(8))
	fmt.Println(fact(9))
	fmt.Println(fact(10))

}
