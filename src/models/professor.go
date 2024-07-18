package models

type Professor struct {
	Id     uint    `gorm:"primaryKey"`
	Nome   string  `gorm:"type:varchar(255)"`
	Email  string  `gorm:"type:varchar(255)"`
	CPF    string  `gorm:"type:varchar(255)"`
	Turmas []Turma `gorm:"foreignKey:ProfessorId"`
}

func (Professor) TableName() string {
	return "professor"
}
