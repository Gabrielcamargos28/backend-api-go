package data

type AlunoResponse struct {
	Id        uint            `json:"id"`
	Nome      string          `json:"nome"`
	Matricula string          `json:"matricula"`
	Turmas    []TurmaResponse `json:"turmas"`
	Notas     []AlunoNota     `json:"notas"`
}
