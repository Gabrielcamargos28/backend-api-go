package aluno

import (
	"controle-notas/src/data/aluno/request"
	"controle-notas/src/data/aluno/response"
)

type AlunoService interface {
	Create(aluno request.AlunoRequest)
	Update(aluno request.AtualizarAlunoRequest)
	Delete(alunoId uint)
	FindById(aluno uint) response.AlunoResponse
	FindAll() []response.AlunoResponse
}
