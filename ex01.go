package main

import (
	"fmt"
)

func encontrarMaiorAlgarismo(lista []int) int {
	maior := 0

	for _, num := range lista {
		for num > 0 {
			algarismo := num % 10
			if algarismo > maior {
				maior = algarismo
			}
			num /= 10
		}
	}

	return maior
}

func main() {
	numeros := []int{123, 456, 7890, 98765, 4321}
	maiorAlgarismo := encontrarMaiorAlgarismo(numeros)
	fmt.Println("Maior algarismo:", maiorAlgarismo)
}
