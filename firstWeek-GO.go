package main

import (
	"fmt"
	"math"
)

// Calcula o discriminante (delta) de uma equação do segundo grau
// Retorna delta = b² - 4ac.
func bhaskara(a, b, c int) (int, int) {
	deltaresult := (b * b) - 4*a*c
	SquaretDelta := math.Sqrt(float64(deltaresult)) // Raiz quadrada do delta
	xpos := -b + (int(SquaretDelta))/2*a
	xneg := -b - (int(SquaretDelta))/2*a
	if deltaresult < 0 {
		fmt.Println("Delta é negativo, não existe raiz real.")
	} else if deltaresult == 0 {
		fmt.Println("Delta é igual a zero, existe uma raiz real.")
		return xpos, xneg
	} else {
		fmt.Println("Delta é positivo, existem duas raízes reais.")
		return xpos, xneg
	}
}

func main() {
	fmt.Println(bhaskara(4, 2, -6)) // Exemplo de uso da função delta
	// Resultado esperado é: X' = 18 X'' = -22
	// Delta é positivo, existem duas raízes reais.
}

// O código calcula o discriminante (delta) de uma equação do segundo grau
