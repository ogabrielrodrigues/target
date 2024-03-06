package main

import "fmt"

func inverter(str string) string {
	runes := []rune(str)

	for i, f := 0, len(runes)-1; i < len(runes)/2; i, f = i+1, f-1 {
		runes[i], runes[f] = runes[f], runes[i]
	}

	return string(runes)
}

func main() {
	fmt.Println(inverter("meu nome é gabriel"))
}
