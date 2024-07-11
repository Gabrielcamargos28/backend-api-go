package request

import (
	atividade "controle-notas/src/controller/model/request"
	"controle-notas/src/controller/model/request/aluno"
)

type Nota struct {
	ID          uint                `json:"id" gorm:"primary_key"`
	Aluno       aluno.Aluno         `json:"aluno" gorm:"foreignkey:AlunoID"`
	AlunoID     uint                `json:"alunoId"`
	Atividade   atividade.Atividade `json:"atividade" gorm:"foreignkey:AtividadeID"`
	AtividadeID uint                `json:"atividadeId"`
	Valor       float64             `json:"valor"`
}
