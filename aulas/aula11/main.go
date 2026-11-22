package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		mostrarUso()
		return
	}

	tarefas, err := carregarTarefas()
	if err != nil {
		fmt.Println("Erro ao carregar tarefas:", err)
		return
	}

	comando := os.Args[1]

	switch comando {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Informe o título. Exemplo: tarefas add Estudar Go")
			return
		}
		titulo := strings.Join(os.Args[2:], " ")
		tarefas = adicionar(tarefas, titulo)

	case "list":
		listar(tarefas)
		return

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Informe o número da tarefa. Exemplo: tarefas done 1")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Número inválido:", os.Args[2])
			return
		}
		err = concluir(tarefas, id)
		if err != nil {
			fmt.Println(err)
			return
		}

	default:
		mostrarUso()
		return
	}

	err = salvarTarefas(tarefas)
	if err != nil {
		fmt.Println("Erro ao salvar tarefas:", err)
	}
}

func mostrarUso() {
	fmt.Println("Como usar:")
	fmt.Println("  tarefas add <título>   adiciona uma tarefa")
	fmt.Println("  tarefas list           mostra todas as tarefas")
	fmt.Println("  tarefas done <número>  marca a tarefa como feita")
}

func adicionar(tarefas []Tarefa, titulo string) []Tarefa {
	proximoID := 1
	if len(tarefas) > 0 {
		ultima := tarefas[len(tarefas)-1]
		proximoID = ultima.ID + 1
	}

	nova := Tarefa{
		ID:       proximoID,
		Titulo:   titulo,
		CriadaEm: time.Now(),
	}
	fmt.Println("Tarefa adicionada:", nova.ID, "-", nova.Titulo)
	return append(tarefas, nova)
}

func listar(tarefas []Tarefa) {
	if len(tarefas) == 0 {
		fmt.Println("Nenhuma tarefa ainda. Use: tarefas add <título>")
		return
	}
	for _, t := range tarefas {
		status := "[ ]"
		if t.Concluida {
			status = "[x]"
		}
		fmt.Println(status, t.ID, "-", t.Titulo)
	}
}

func concluir(tarefas []Tarefa, id int) error {
	for i := range tarefas {
		if tarefas[i].ID == id {
			tarefas[i].Concluida = true
			fmt.Println("Tarefa", id, "concluída!")
			return nil
		}
	}
	return fmt.Errorf("tarefa %d não encontrada", id)
}
