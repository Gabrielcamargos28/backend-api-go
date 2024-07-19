package repository

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/models"
)

type NotaRepository interface {
	Save(nota models.Nota) *rest_err.RestErr
	Update(nota models.Nota) *rest_err.RestErr
	Delete(notaId uint) *rest_err.RestErr
	FindById(notaId uint) (models.Nota, *rest_err.RestErr)
	FindAll() ([]models.Nota, *rest_err.RestErr)
	FindNotasByAlunoId(alunoId uint) ([]models.Nota, *rest_err.RestErr)
	FindByAlunoAndAtividade(alunoId uint, atividadeId uint) (*models.Nota, *rest_err.RestErr)
}
