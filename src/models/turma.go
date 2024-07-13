package models

type Turma struct {
	Id          uint        `gorm:"primaryKey"`
	Nome        string      `gorm:"type:varchar(255);not null"`
	Semestre    string      `gorm:"type:varchar(50);not null"`
	Ano         int         `gorm:"not null"`
	ProfessorId uint        `gorm:"not null"`
	Professor   Professor   `gorm:"foreignKey:ProfessorId"`
	Atividades  []Atividade `gorm:"foreignKey:TurmaID"`
	Alunos      []Aluno     `gorm:"many2many:turma_alunos;"`
}
