package data

type AtualizarNota struct {
	Id    uint    `json:"id" validate:"required"`
	Valor float64 `json:"valor" binding:"required"`
}
