package repository

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/models"
)

type TurmaRepository interface {
	Save(turma models.Turma) *rest_err.RestErr
	Update(turma models.Turma) *rest_err.RestErr
	Delete(turmaId uint) *rest_err.RestErr
	FindById(turmaId uint) (models.Turma, *rest_err.RestErr)
	FindAll() ([]models.Turma, *rest_err.RestErr)
	RemoveAlunoTurma(turmaId uint, alunoId uint) *rest_err.RestErr
	FindAtividadesByTurmaId(turmaId uint) ([]models.Atividade, *rest_err.RestErr)
}
