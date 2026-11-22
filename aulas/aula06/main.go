package main

import (
	"errors"
	"fmt"
)

func somar(a int, b int) int {
	return a + b
}

func saudacao(nome string) string {
	return "Olá, " + nome + "!"
}

func dividir(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("não dá para dividir por zero")
	}
	return a / b, nil
}

func processarPedido(valor float64, estoque int, pago bool) string {
	if valor <= 0 {
		return "Erro: valor inválido"
	}
	if estoque <= 0 {
		return "Erro: sem estoque"
	}
	if !pago {
		return "Erro: pagamento pendente"
	}

	return "Pedido enviado!"
}

var ErrSaldoInsuficiente = errors.New("saldo insuficiente")

func sacar(saldo float64, valor float64) (float64, error) {
	if valor > saldo {
		return saldo, ErrSaldoInsuficiente
	}
	return saldo - valor, nil
}

func processarSaque(conta string, saldo float64, valor float64) error {
	_, err := sacar(saldo, valor)
	if err != nil {
		return fmt.Errorf("saque na conta %s: %w", conta, err)
	}
	return nil
}

func dobrar(x int) int {
	dobro := x * 2
	return dobro
}

func main() {
	total := somar(2, 3)
	fmt.Println(total)
	fmt.Println(saudacao("Ana"))

	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Resultado:", resultado)

	_, err = dividir(10, 0)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Println(processarPedido(150, 3, true))
	fmt.Println(processarPedido(150, 0, true))

	err = processarSaque("0042", 100, 250)
	fmt.Println(err)

	if errors.Is(err, ErrSaldoInsuficiente) {
		fmt.Println("→ Oferecer limite especial ao cliente")
	}

	fmt.Println(dobrar(21))
}
