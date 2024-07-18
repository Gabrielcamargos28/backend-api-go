package turma

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data"
)

type TurmaService interface {
	Create(turma data.TurmaRequest) *rest_err.RestErr
	Update(turma data.AtualizaTurmaRequest) *rest_err.RestErr
	Delete(turmaId uint) *rest_err.RestErr
	FindById(turmaId uint) (data.TurmaResponse, *rest_err.RestErr)
	FindAll() ([]data.TurmaResponse, *rest_err.RestErr)
	AdicionarAlunos(turma data.AdicioanarAlunosTurma) *rest_err.RestErr
	RemoveAlunoTurma(alunoId uint, turmaId uint) *rest_err.RestErr
}
