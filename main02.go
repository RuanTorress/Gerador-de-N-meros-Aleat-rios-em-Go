package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var quantity, min, max int

	// Definir uma semente para o gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Gerador de Números Aleatórios")
	fmt.Println("------------------------------")

	// Solicitar ao usuário a quantidade de números a gerar
	fmt.Print("Quantos números aleatórios você deseja gerar? ")
	fmt.Scanln(&quantity)

	// Solicitar ao usuário o intervalo de números
	fmt.Print("Digite o valor mínimo: ")
	fmt.Scanln(&min)
	fmt.Print("Digite o valor máximo: ")
	fmt.Scanln(&max)

	// Validar o intervalo
	if min >= max {
		fmt.Println("Erro: O valor mínimo deve ser menor que o valor máximo.")
		return
	}

	// Gerar e exibir os números aleatórios
	fmt.Printf("Números aleatórios gerados entre %d e %d:\n", min, max)
	for i := 0; i < quantity; i++ {
		randomNumber := rand.Intn(max-min+1) + min
		fmt.Println(randomNumber)
	}
}
