package data

type AtualizarNota struct {
	NotaId float64 `json:"id" binding:"required"`
	Valor  float64 `json:"valor" binding:"required"`
}
