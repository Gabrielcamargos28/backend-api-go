package data

type TurmaAlunosResponse struct {
	Id        uint            `json:"id"`
	Nome      string          `json:"nome"`
	Semestre  string          `json:"semestre"`
	Ano       int             `json:"ano"`
	Professor string          `json:"professor"`
	Alunos    []AlunoResponse `json:"alunos"`
}
