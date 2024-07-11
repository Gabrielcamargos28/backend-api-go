package atividade

type Atividade struct {
	ID      uint    `json:"id" gorm:"primary_key"`
	Turma   Turma   `json:"turma" gorm:"foreignkey:TurmaID"`
	TurmaID uint    `json:"turmaId"`
	Valor   float64 `json:"valor"`
	Data    string  `json:"data"`
}
