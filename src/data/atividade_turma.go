package data

import (
	"time"
)

type AtividadeTurmaResponse struct {
	Id    uint      `json:"id"`
	Nome  string    `json:"nome"`
	Valor float64   `json:"valor"`
	Data  time.Time `json:"data"`
}
