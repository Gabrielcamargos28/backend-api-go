package atividade

import "controle-notas/src/controller/model/request/turma"

type Atividade struct {
	ID      uint        `json:"id" gorm:"primary_key"`
	Turma   turma.Turma `json:"turma" gorm:"foreignkey:TurmaID"`
	TurmaID uint        `json:"turmaId"`
	Valor   float64     `json:"valor"`
	Data    string      `json:"data"`
}
