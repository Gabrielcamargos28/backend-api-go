package data

type AlunoNota struct {
	AlunoId       uint    `json:"alunoId"`
	AlunoNome     string  `json:"alunoNome"`
	Nota          float64 `json:"nota"`
	TurmaId       uint    `json:"turmaId"`
	TurmaNome     string  `json:"turmaNome"`
	AtividadeId   uint    `json:"atividadeId"`
	AtividadeNome string  `json:"atividadeNome"`
}
