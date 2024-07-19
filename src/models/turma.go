package models

type Turma struct {
	Id          uint        `gorm:"primaryKey"`
	Nome        string      `gorm:"type:varchar(255);not null"`
	Semestre    string      `gorm:"type:varchar(50);not null"`
	Ano         int         `gorm:"not null"`
	ProfessorId uint        `gorm:"index"`
	Professor   Professor   `gorm:"contraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Atividades  []Atividade `gorm:"foreignKey:TurmaId"`
	Alunos      []Aluno     `gorm:"many2many:turma_alunos;constraint:OnDelete:CASCADE;"`
}

func (Turma) TableName() string {
	return "turma"
}
