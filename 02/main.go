package main

import (
	"fmt"
	"math"
)

/*
	Fórmula de Binet

	Um número é Fibonacci se e apenas se uma
	ou ambas as expressões 5 * n^2 + 4 ou
	5 * n^2 – 4 retornarem um quadrado perfeito.
*/

func calcular(x int) bool {
	y := math.Sqrt(float64(x))

	return y*y == float64(x)
}

func fibonacci(numero int) bool {
	return calcular(5*numero*numero-4) || calcular(5*numero*numero+4)
}

func main() {
	var n = 256

	if fibonacci(n) {
		fmt.Printf("%d faz parte da sequência de fibonacci.\n", n)
	} else {
		fmt.Printf("%d não faz parte da sequência de fibonacci.\n", n)
	}
}
