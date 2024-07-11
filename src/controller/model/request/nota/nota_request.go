package request

type Nota struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Aluno       Aluno     `json:"aluno" gorm:"foreignkey:AlunoID"`
	AlunoID     uint      `json:"alunoId"`
	Atividade   Atividade `json:"atividade" gorm:"foreignkey:AtividadeID"`
	AtividadeID uint      `json:"atividadeId"`
	Valor       float64   `json:"valor"`
}
