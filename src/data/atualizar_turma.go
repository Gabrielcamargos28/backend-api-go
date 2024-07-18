package data

type AtualizaTurmaRequest struct {
	Id          uint   `json:"id"`
	Nome        string `json:"nome"`
	Semestre    string `json:"semestre"`
	Ano         int    `json:"ano"`
	ProfessorId uint   `json:"professor_id"`
}
