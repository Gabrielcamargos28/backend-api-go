package data

type TurmaAlunosResponse struct {
	Id             uint                     `json:"id"`
	Nome           string                   `json:"nome"`
	Semestre       string                   `json:"semestre"`
	Ano            int                      `json:"ano"`
	ProfessorId    uint                     `json:"professorId"`
	Professor      string                   `json:"professor"`
	Alunos         []AlunoResumido          `json:"alunos"`
	Atividades     []AtividadeTurmaResponse `json:"atividades"`
	SomaAtividades float64                  `json:"somaAtividades"`
}
