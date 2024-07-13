package models

import (
	"time"
)

type Atividade struct {
	ID      uint      `gorm:"primaryKey"`
	Valor   float64   `gorm:"not null"`
	Data    time.Time `gorm:"not null"`
	TurmaID uint      `gorm:"not null"`
	Turma   Turma     `gorm:"foreignKey:TurmaID"`
	Notas   []Nota    `gorm:"foreignKey:AtividadeID"`
}

func (Atividade) TableName() string {
	return "atividade"
}
