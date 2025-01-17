// Go supports _constants_ of character, string, boolean,
// and numeric values.

package main

import "fmt"
import "math"

// `const` declara um valor constante.
const s string = "constante"

func main() {
    fmt.Println(s)
    
    // O estatamento `const` pode aparecer em qualquer lugar 
    // onde o estatamento `var` pode.    
    const n = 500000000

    // Expressões constantes realizam aritmética com
    // precisão arbitrária.    
    const d = 3e20 / n
	fmt.Println(3e20)
    fmt.Println(d)

    // Uma constante numérica não possui nenhum tipo até
    // que seja dado um, como por uma conversão explícita.    
    fmt.Println(int64(d))
    
    // Um número pode ser definido como um tipo usando ele
    // em um contexto que requira um, assim como uma
    // atribuição de variável ou função chamada. Por
    // exemplo, aqui `math.Sin` espera um `float64`.
    fmt.Println(math.Sin(n))
}