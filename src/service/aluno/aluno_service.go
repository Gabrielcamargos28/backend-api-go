package aluno

import (
	"controle-notas/src/data"
)

type AlunoService interface {
	Create(aluno data.AlunoRequest)
	Update(aluno data.AtualizarAlunoRequest)
	Delete(alunoId uint)
	FindById(aluno uint) data.AlunoResponse
	FindAll() []data.AlunoResponse
}
