package professor

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data"
)

type ProfessorService interface {
	Create(professor data.ProfessorRequest) *rest_err.RestErr
	Update(professor data.AtualizarProfessorRequest) *rest_err.RestErr
	Delete(professorId uint) *rest_err.RestErr
	FindById(professorId uint) (data.ProfessorResponse, *rest_err.RestErr)
	FindAll() ([]data.ProfessorResponse, *rest_err.RestErr)
}
