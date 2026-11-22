package main

import "fmt"

func main() {
	idade := 17

	if idade >= 18 {
		fmt.Println("Pode tirar carteira de motorista")
	} else {
		fmt.Println("Ainda não pode tirar carteira")
	}

	temperatura := 38.5

	if temperatura >= 37.8 {
		fmt.Println("Febre")
	} else if temperatura >= 37.0 {
		fmt.Println("Atenção")
	} else {
		fmt.Println("Normal")
	}

	temChave := true
	temSenha := false

	if temChave && temSenha {
		fmt.Println("Acesso liberado")
	} else {
		fmt.Println("Acesso negado")
	}

	if temChave || temSenha {
		fmt.Println("Tem pelo menos uma das duas")
	}

	if !temSenha {
		fmt.Println("Falta a senha")
	}

	preco := 500.0

	if desconto := preco * 0.15; desconto > 50 {
		fmt.Println("Desconto alto:", desconto)
	}

	dia := "sábado"

	switch dia {
	case "sábado", "domingo":
		fmt.Println("Fim de semana")
	case "sexta":
		fmt.Println("Quase lá")
	default:
		fmt.Println("Dia útil")
	}

	codigo := 503

	switch {
	case codigo >= 200 && codigo < 300:
		fmt.Println("Sucesso")
	case codigo >= 400 && codigo < 500:
		fmt.Println("Erro de quem pediu")
	case codigo >= 500:
		fmt.Println("Erro no servidor")
	default:
		fmt.Println("Outro")
	}
}
