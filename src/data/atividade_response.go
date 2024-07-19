package data

import (
	"time"
)

type AtividadeResponse struct {
	Id    uint          `json:"id"`
	Nome  string        `json:"nome"`
	Valor float64       `json:"valor"`
	Data  time.Time     `json:"data"`
	Turma TurmaResponse `json:"turma"`
	Notas []AlunoNota   `json:"alunos_notas"`
}
