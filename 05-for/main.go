// `for` é a única estrutura de looping do Go. Aqui estão
// três tipos básicos de `for` loops.

package main

import "fmt"

func main() {

	// O tipo mais comum, com uma única condição.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Uma clássica inicial/condição/final `for` loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` sem uma condição se repetirá várias vezes
	// até que você dê um `break` no loop ou `return` da
	// função de fechamento.
	for {
		fmt.Println("loop")
		break
	}
}