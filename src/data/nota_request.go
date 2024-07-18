package data

type NotaRequest struct {
	Valor       float64 `json:"valor" binding:"required"`
	AlunoId     uint    `json:"aluno_id" binding:"required"`
	AtividadeId uint    `json:"atividade_id" binding:"required"`
}
