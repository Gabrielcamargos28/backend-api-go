package repository

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/models"
)

type AtividadeRepository interface {
	Save(atividade models.Atividade) *rest_err.RestErr
	Update(atividade models.Atividade) *rest_err.RestErr
	Delete(atividadeId uint) *rest_err.RestErr
	FindById(atividadeId uint) (models.Atividade, *rest_err.RestErr)
	FindAll() ([]models.Atividade, *rest_err.RestErr)
}
