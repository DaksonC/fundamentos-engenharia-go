package main

import "time"

type Tarefa struct {
	ID        int       `json:"id"`
	Titulo    string    `json:"titulo"`
	Concluida bool      `json:"concluida"`
	CriadaEm  time.Time `json:"criada_em"`
}
