package data

type NotaRequest struct {
	Valor       float64 `json:"valor" binding:"required"`
	AlunoId     uint    `json:"alunoId" binding:"required"`
	AtividadeId uint    `json:"atividadeId" binding:"required"`
}
