package turma

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data/turma/request"
	"controle-notas/src/data/turma/response"
)

type TurmaService interface {
	Create(turma request.TurmaRequest) *rest_err.RestErr
	Update(turma request.AtualizaTurmaRequest) *rest_err.RestErr
	Delete(turmaId uint) *rest_err.RestErr
	FindById(turmaId uint) (response.TurmaResponse, *rest_err.RestErr)
	FindAll() ([]response.TurmaResponse, *rest_err.RestErr)
	AdicionarAlunos(turma request.AdicioanrAlunosTurma) *rest_err.RestErr
	RemoveAlunoTurma(alunoId uint, turmaId uint) *rest_err.RestErr
}
