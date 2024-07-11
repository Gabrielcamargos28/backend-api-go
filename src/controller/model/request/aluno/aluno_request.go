package aluno

import "controle-notas/src/controller/model/request/turma"

type Aluno struct {
	ID        uint          `json:"id" gorm:"primary_key"`
	Nome      string        `json:"nome"`
	Matricula string        `json:"matricula"`
	Turmas    []turma.Turma `json:"turmas" gorm:"many2many:aluno_turmas;"`
}
