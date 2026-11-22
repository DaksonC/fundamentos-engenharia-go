package main

import "fmt"

func tentarZerar(valor int) {
	valor = 0
}

func zerarDeVerdade(valor *int) {
	*valor = 0
}

func main() {
	saldo := 10000 // R$ 100,00 em centavos
	fmt.Println("valor:", saldo)
	fmt.Println("endereço:", &saldo)

	ponteiro := &saldo
	fmt.Printf("ponteiro guarda: %v | tipo: %T\n", ponteiro, ponteiro)

	fmt.Println("indo até o endereço:", *ponteiro)

	*ponteiro = 25000
	fmt.Println("saldo agora:", saldo)

	var p *int
	fmt.Println(p)

	if p != nil {
		fmt.Println(*p)
	}

	conta := 50000 // R$ 500,00 em centavos

	tentarZerar(conta)
	fmt.Println("depois de tentarZerar:", conta)

	zerarDeVerdade(&conta)
	fmt.Println("depois de zerarDeVerdade:", conta)

	cv := configPadraoValor()
	cp := configPadraoPonteiro()
	fmt.Println(cv.Porta, cp.Porta)
}
