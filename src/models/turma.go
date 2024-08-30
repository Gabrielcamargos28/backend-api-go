package models

import "time"

type Turma struct {
	Id          uint        `gorm:"primaryKey"`
	Nome        string      `gorm:"type:varchar(255);not null"`
	Semestre    string      `gorm:"type:varchar(50);not null"`
	Ano         int         `gorm:"not null"`
	ProfessorId uint        `gorm:"index"`
	Professor   Professor   `gorm:"constraint:OnUpdate;"`
	Atividades  []Atividade `gorm:"foreignKey:TurmaId;"`
	Alunos      []Aluno     `gorm:"many2many:turma_alunos;"`
	DeletedAt   *time.Time  `gorm:"index"`
}

func (Turma) TableName() string {
	return "turma"
}
