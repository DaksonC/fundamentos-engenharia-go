package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Forma completa: nome, tipo e valor
	var nome string = "Cruz"

	// Forma curta: o Go descobre o tipo sozinho
	idade := 43
	salarioHora := 120.5
	ativo := true

	fmt.Println(nome, idade, salarioHora, ativo)
	fmt.Printf("nome é %T, idade é %T, salarioHora é %T, ativo é %T\n",
		nome, idade, salarioHora, ativo)

	var contador int
	var preco float64
	var logado bool
	var mensagem string

	fmt.Println("contador:", contador)
	fmt.Println("preco:", preco)
	fmt.Println("logado:", logado)
	fmt.Printf("mensagem: %q\n", mensagem)

	const taxaServico = 0.10
	fmt.Println("Taxa:", taxaServico)

	fmt.Println("bool:   ", unsafe.Sizeof(logado), "byte")
	fmt.Println("int:    ", unsafe.Sizeof(contador), "bytes")
	fmt.Println("float64:", unsafe.Sizeof(preco), "bytes")
	fmt.Println("string: ", unsafe.Sizeof(mensagem), "bytes")

	a := 0.1
	b := 0.2
	fmt.Println(a + b)
}
