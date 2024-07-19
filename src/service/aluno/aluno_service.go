package aluno

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data"
)

type AlunoService interface {
	Create(aluno data.AlunoRequest) *rest_err.RestErr
	Update(aluno data.AtualizarAlunoRequest) *rest_err.RestErr
	Delete(alunoId uint) *rest_err.RestErr
	FindById(alunoId uint) (data.AlunoResponse, *rest_err.RestErr)
	FindAll() ([]data.AlunoResponse, *rest_err.RestErr)
}
