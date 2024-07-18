package data

import "time"

type AtividadeRequest struct {
	Id      uint      `json:"id" validade:required`
	Nome    string    `json:"nome" validade:required`
	Valor   float64   `json:"valor" validate:"required"`
	Data    time.Time `json:"data" validate:"required"`
	TurmaId uint      `json:"turma_id" validate:"required"`
}