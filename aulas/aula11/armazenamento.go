package main

import (
	"encoding/json"
	"errors"
	"os"
)

const arquivo = "tarefas.json"

func carregarTarefas() ([]Tarefa, error) {
	dados, err := os.ReadFile(arquivo)
	if errors.Is(err, os.ErrNotExist) {
		return []Tarefa{}, nil
	}
	if err != nil {
		return nil, err
	}

	var tarefas []Tarefa
	err = json.Unmarshal(dados, &tarefas)
	if err != nil {
		return nil, err
	}
	return tarefas, nil
}

func salvarTarefas(tarefas []Tarefa) error {
	dados, err := json.MarshalIndent(tarefas, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(arquivo, dados, 0644)
}
