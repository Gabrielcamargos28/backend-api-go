package models

type Professor struct {
	Id     uint    `gorm:"primary_key`
	Nome   string  `gorm:type:varchar(255)`
	Email  string  `gorm:type:varchar(255)`
	CPF    string  `gorm:type:varchar(255)`
	Turmas []Turma `gorm:"foreignKey:ProfessorID"`
}

func (Professor) TableName() string {
	return "professor"
}
