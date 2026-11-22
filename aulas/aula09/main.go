package main

import (
	"errors"
	"fmt"
)

type Pedido struct {
	ID          int
	NomeCliente string
	Valor       int
	Pago        bool
}

func NovoPedido(id int, cliente string, valor int) (*Pedido, error) {
	if valor <= 0 {
		return nil, errors.New("o valor do pedido precisa ser maior que zero")
	}
	return &Pedido{ID: id, NomeCliente: cliente, Valor: valor}, nil
}

func (p *Pedido) Resumo() string {
	status := "pendente"
	if p.Pago {
		status = "pago"
	}
	return fmt.Sprintf("Pedido %d de %s: R$ %d,%02d (%s)",
		p.ID, p.NomeCliente, p.Valor/100, p.Valor%100, status)
}

func (p *Pedido) ConfirmarPagamento() {
	p.Pago = true
}

type Endereco struct {
	Rua    string
	Cidade string
}

func (e Endereco) Completo() string {
	return e.Rua + ", " + e.Cidade
}

type Cliente struct {
	Nome string
	Endereco
}

func main() {
	p := Pedido{
		ID:          1,
		NomeCliente: "Ana",
		Valor:       8990,
	}

	fmt.Println(p.NomeCliente, p.Valor, p.Pago)
	fmt.Printf("%+v\n", p)

	pedido, err := NovoPedido(2, "Bruno", 15000)
	if err != nil {
		fmt.Println(err)
		return
	}
	pedido.ConfirmarPagamento()
	fmt.Println(pedido.Resumo())

	c := Cliente{
		Nome:     "Ana",
		Endereco: Endereco{Rua: "Rua das Flores, 10", Cidade: "Florianópolis"},
	}
	fmt.Println(c.Cidade)
	fmt.Println(c.Completo())
}
