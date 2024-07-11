package model

import (
	"controle-notas/src/controller/model/request/professor"
)

type Turma struct {
	ID          uint                `json:"id" gorm:"primary_key"`
	Nome        string              `json:"nome"`
	Semestre    string              `json:"semestre"`
	Ano         int                 `json:"ano"`
	Professor   professor.Professor `json:"professor" gorm:"foreignkey:ProfessorID"`
	ProfessorID uint                `json:"professorId"`
}
