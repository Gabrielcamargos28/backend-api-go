package data

type AtualizarNota struct {
	Id    uint    `json:"id" binding:"required"`
	Valor float64 `json:"valor" binding:"required"`
}
