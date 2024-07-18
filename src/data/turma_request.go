package data

type TurmaRequest struct {
	Nome        string `json:"nome"`
	Semestre    string `json:"semestre"`
	Ano         int    `json:"ano"`
	ProfessorId uint   `json:"professor_id"`
}
