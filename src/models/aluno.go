package models

type Aluno struct {
	ID        uint    `gorm:"primaryKey"`
	Nome      string  `gorm:"type:varchar(255);not null"`
	Matricula string  `gorm:"type:varchar(255);unique;not null"`
	Turmas    []Turma `gorm:"many2many:turma_alunos;"`
}

func (Aluno) TableName() string {
	return "aluno"
}
