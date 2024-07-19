package data

import "time"

type NotaResponse struct {
	Id             uint      `json:"id"`
	Valor          float64   `json:"valor"`
	AtividadeId    uint      `json:"atividadeId"`
	Atividade      string    `json:"atividade"`
	AtividadeValor uint      `json:"atividadeValor"`
	Data           time.Time `json:"data"`
}
