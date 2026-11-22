package main

import (
	"fmt"
	"time"
)

func main() {
	// Forma 1: com contador
	for i := 1; i <= 5; i++ {
		fmt.Println("Volta", i)
	}

	// Forma 2: só a condição
	bateria := 100
	for bateria > 0 {
		fmt.Println("Bateria:", bateria)
		bateria = bateria - 30
	}

	// Forma 3: um número de vezes
	for i := range 3 {
		fmt.Println("Tentativa", i)
	}

	// Forma 4: para sempre, até um break
	ticks := 0
	for {
		ticks++
		fmt.Println("Verificando se chegou trabalho novo... vez", ticks)
		time.Sleep(500 * time.Millisecond)

		if ticks == 5 {
			fmt.Println("Desligando.")
			break
		}
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i > 7 {
			break
		}
		fmt.Println("Ímpar:", i)
	}

	n := 1_000
	comparacoes := 0
	inicio := time.Now()

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			comparacoes++
		}
	}

	fmt.Println("Itens:", n)
	fmt.Println("Comparações:", comparacoes)
	fmt.Println("Tempo:", time.Since(inicio))
}
