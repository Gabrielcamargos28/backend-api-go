package repository

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/models"
)

type ProfessorRepository interface {
	Save(professor models.Professor) *rest_err.RestErr
	Update(professor models.Professor) *rest_err.RestErr
	Delete(professorId uint) *rest_err.RestErr
	FindById(professorId uint) (models.Professor, *rest_err.RestErr)
	FindAll() ([]models.Professor, *rest_err.RestErr)
}
