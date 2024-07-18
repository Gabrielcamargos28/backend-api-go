package models

import (
	"time"
)

type AlunoNota struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	AtividadeID uint      `json:"atividade_id"`
	AlunoID     uint      `json:"aluno_id"`
	Nota        float64   `json:"nota"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (AlunoNota) TableName() string {
	return "alunos_notas"
}
