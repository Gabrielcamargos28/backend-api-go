package repository

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/models"
)

type AlunoRepository interface {
	Save(aluno models.Aluno) *rest_err.RestErr
	Update(aluno models.Aluno) *rest_err.RestErr
	Delete(alunoId uint) *rest_err.RestErr
	FindById(alunoId uint) (models.Aluno, *rest_err.RestErr)
	FindAll() ([]models.Aluno, *rest_err.RestErr)
	FindNotasByAlunoId(alunoId uint) ([]models.Nota, *rest_err.RestErr)
}
