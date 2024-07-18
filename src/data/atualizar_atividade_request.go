package data

import "time"

type AtualizarAtividadeRequest struct {
	Id        uint      `json:"id" validate:"required"`
	Nome      string    `json:"nome" validate:"required"`
	TurmaId   uint      `json:"turma_id" validate:"required"`
	TurmaNome string    `json:"turma_nome" validate:"required"`
	Valor     float64   `json:"valor" validate:"required"`
	Data      time.Time `json:"data" validate:"required"`
}
