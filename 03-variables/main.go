// No Go, _variáveis_ são explicitamente declaradas e usadas pelo
// compilador para, por exemplo, checar o tipo-correção das
// funções chamadas.

package main

import "fmt"

func main() {

    // `var` declara 1 ou mais variáveis.    
    var a string = "inicial"
    fmt.Println(a)
 
    // Você pode declarar várias variáveis de uma só vez.    
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Go vai inferir o tipo de variáveis inicializadas.    
    var d = true
    fmt.Println(d)

    // Variáveis declaradas sem a inicialização correspondente
    // são _valor-zero_. Por exemplo, o valor zero para
    // `int` é `0`.
    
    var e int
    fmt.Println(e)

    // A sintaxe `:=` é uma abreviação para declarar e
    // inicializar uma variável, e.x. para 
    // `var f string = "curto"` nesse caso.    
    f := "curto"
    fmt.Println(f)
}