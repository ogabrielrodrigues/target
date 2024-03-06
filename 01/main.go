package main

import "fmt"

func main() {
	var indice, soma, k = 13, 0, 0

	// Não existe do-while em Go, porém isso é equivalente:
	for k < indice {
		k = k + 1
		soma = soma + k
	}

	fmt.Printf("SOMA: %d\n", soma)
}
