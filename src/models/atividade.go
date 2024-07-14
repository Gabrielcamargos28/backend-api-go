package models

import (
	"time"
)

type Atividade struct {
	Id      uint      `gorm:"primaryKey"`
	Valor   float64   `gorm:"not null"`
	Data    time.Time `gorm:"not null"`
	TurmaId uint      `gorm:"not null"`
	Turma   Turma     `gorm:"foreignKey:TurmaId"`
	Notas   []Nota    `gorm:"foreignKey:AtividadeId"`
}

func (Atividade) TableName() string {
	return "atividade"
}
