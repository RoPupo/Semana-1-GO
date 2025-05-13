package main

import (
	"fmt"
	"math"
)

// Calcula o discriminante (delta) de uma equação do segundo grau
// Retorna delta = b² - 4ac.
func bhaskara(a, b, c int) (float64, float64, int) {
	delta := (b * b) - 4*a*c
	if delta < 0 { // Se delta for negativo, não existem raízes reais, o código não deve retornar nada, apenas imprimir uma mensagem
		return 0, 0, delta
	}
	SqrtDelta := math.Sqrt(float64(delta)) // Raiz quadrada do delta
	root1 := (-float64(b) + SqrtDelta) / (2 * float64(a))
	root2 := (-float64(b) - SqrtDelta) / (2 * float64(a))

	return root1, root2, delta
}
func main() {
	root1, root2, delta := bhaskara(4, 8, -15)
	if delta < 0 {
		fmt.Println("Delta é negativo, não existem raízes reais para essa equação")
	} else if delta == 0 {
		fmt.Println("Delta é igual a zero, existe apenas uma raiz real para essa equação")
		fmt.Println("Raiz:", root1)
	} else if delta > 0 {
		fmt.Println("Delta é positivo, existem duas raízes reais para essa equação")
		fmt.Printf("Raiz 1: %.4f\n", root1)
		fmt.Printf("Raiz 2:%.4f\n", root2)
	}
}
