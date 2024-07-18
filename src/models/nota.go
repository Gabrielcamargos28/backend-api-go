package models

type Nota struct {
	Id          uint      `gorm:"primaryKey"`
	Valor       float64   `gorm:"not null"`
	AlunoId     uint      `gorm:"not null"`
	Aluno       Aluno     `gorm:"foreignKey:AlunoId"`
	AtividadeId uint      `gorm:"not null"`
	Atividade   Atividade `gorm:"foreignKey:AtividadeId"`
}

func (Nota) TableName() string {
	return "nota"
}
