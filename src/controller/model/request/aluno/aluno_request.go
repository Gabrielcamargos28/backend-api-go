package aluno

type Aluno struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	Nome      string  `json:"nome"`
	Matricula string  `json:"matricula"`
	Turmas    []Turma `json:"turmas" gorm:"many2many:aluno_turmas;"`
}
