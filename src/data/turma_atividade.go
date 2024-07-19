package data

type TurmaAtividades struct {
	Id          uint                     `json:"id"`
	Nome        string                   `json:"nome"`
	Semestre    string                   `json:"semestre"`
	Ano         int                      `json:"ano"`
	ProfessorId uint                     `json:"professorId"`
	Professor   string                   `json:"professor"`
	Alunos      []AlunoResponse          `json:"alunos"`
	Atividades  []AtividadeTurmaResponse `json:"atividades"`
}
