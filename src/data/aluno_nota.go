package data

type AlunoNota struct {
	AlunoId   uint    `json:"aluno_id"`
	AlunoNome string  `json:"aluno_nome"`
	Nota      float64 `json:"nota"`
}
