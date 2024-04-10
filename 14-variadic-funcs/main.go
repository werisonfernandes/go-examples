// [_Funções Variádicas_](http://en.wikipedia.org/wiki/Variadic_function)
// podem ser chamadas com qualquer número de argumentos à direta.
// Por exemplo, `fmt.Println` é uma função variádica comum.

package main

import "fmt"

// Aqui temos uma função que irá assumir um número arbitrário
// de `int`eiros como argumentos.
func soma(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// Funções variádicas podem ser chamadas da forma usual
	// com argumentos individuais.
	soma(1, 2)
	soma(1, 2, 3)

	// Se você possui múltiplos argumentos em um slice,
	// aplique-os em uma função variádica usando
	// `func(slice...)` assim.
	nums := []int{1, 2, 3, 4}
	soma(nums...)
}
