package models

type Nota struct {
	Id          uint      `gorm:"primaryKey"`
	Valor       float64   `gorm:"not null"`
	AlunoID     uint      `gorm:"not null"`
	Aluno       Aluno     `gorm:"foreignKey:AlunoID"`
	AtividadeID uint      `gorm:"not null"`
	Atividade   Atividade `gorm:"foreignKey:AtividadeID"`
}

func (Nota) TableName() string {
	return "nota"
}
