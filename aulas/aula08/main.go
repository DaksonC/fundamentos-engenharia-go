package main

import (
	"fmt"
	"time"
)

func temDuplicadoLento(ids []int) bool {
	for i := 0; i < len(ids); i++ {
		for j := i + 1; j < len(ids); j++ {
			if ids[i] == ids[j] {
				return true
			}
		}
	}
	return false
}

func temDuplicadoRapido(ids []int) bool {
	vistos := make(map[int]bool)
	for _, id := range ids {
		if vistos[id] {
			return true
		}
		vistos[id] = true
	}
	return false
}

func main() {
	contas := make(map[string]int)

	contas["ana"] = 150000
	contas["bruno"] = 32050
	contas["ana"] = 175000
	delete(contas, "bruno")

	fmt.Println(contas, "tamanho:", len(contas))

	portas := map[string]int{
		"http":  80,
		"https": 443,
		"ssh":   22,
	}
	fmt.Println("https usa a porta", portas["https"])

	contas["carla"] = 0

	fmt.Println(contas["carla"])
	fmt.Println(contas["joao"])

	saldo, existe := contas["joao"]
	if !existe {
		fmt.Println("Conta joao não encontrada")
	} else {
		fmt.Println("Saldo em centavos:", saldo)
	}

	for protocolo, porta := range portas {
		fmt.Println(protocolo, porta)
	}

	ids := make([]int, 100_000)
	for i := range ids {
		ids[i] = i
	}

	inicio := time.Now()
	temDuplicadoLento(ids)
	fmt.Println("Dois for:", time.Since(inicio))

	inicio = time.Now()
	temDuplicadoRapido(ids)
	fmt.Println("Com map: ", time.Since(inicio))
}
