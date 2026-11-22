package main

import (
	"fmt"
	"time"
)

func semPreAlocar(n int) {
	inicio := time.Now()
	var s []int
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
	fmt.Println("sem make:", time.Since(inicio))
}

func comPreAlocar(n int) {
	inicio := time.Now()
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
	fmt.Println("com make:", time.Since(inicio))
}

func main() {
	var notas [4]int
	notas[0] = 10
	notas[1] = 8
	fmt.Println(notas, "tamanho:", len(notas))

	copia := notas
	copia[0] = 0
	fmt.Println("original:", notas, "cópia:", copia)

	frutas := []string{"maçã", "banana"}
	fmt.Println(frutas[0])

	frutas = append(frutas, "uva")
	fmt.Println(frutas, "tamanho:", len(frutas))

	var pedidos []int
	capAnterior := cap(pedidos)

	for i := 0; i < 2000; i++ {
		pedidos = append(pedidos, i)
		if cap(pedidos) != capAnterior {
			fmt.Println("len:", len(pedidos), "cap:", cap(pedidos), "← realocou!")
			capAnterior = cap(pedidos)
		}
	}

	fila := []string{"Ana", "Bruno", "Carla", "Davi"}
	primeiros := fila[0:2]

	primeiros[0] = "ALTERADO"
	fmt.Println("fila:", fila)

	independente := make([]string, 2)
	copy(independente, fila[0:2])
	independente[0] = "Só aqui"
	fmt.Println("fila:", fila, "| independente:", independente)

	semPreAlocar(10_000_000)
	comPreAlocar(10_000_000)
}
