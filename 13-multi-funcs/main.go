// Go foi construído para suportar _múltiplos valores de retorno_.
// Esse recurso é usado muitas vezes na idiomática do Go, por exemplo
// para retornar ambos os resultados e valores de erros de uma função.

package main

import "fmt"

// O `(int, int)` nessa assinatura de função mostra que
// a função retorna 2 `int`eiros.
func vals() (int, int) {
	return 3, 7
}

func vals2() (int, int, int) {
	return 3, 7, 9
}

func main() {

	// Aqui nós usamos 2 valores de retorno diferentes da
	// chamada com _múltiplas atribuições_.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Se você quer apenas um subconjunto de valores de retorno,
	// use o identificador em branco `_`.
	_, c := vals()
	fmt.Println(c)

	var _, _, result = vals2()
	fmt.Println(result)
}

// todo: parâmetros de retorno nomeados
// todo: retornos crús
