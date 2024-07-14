package professor

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data/professor/request"
	"controle-notas/src/data/professor/response"
)

type ProfessorService interface {
	Create(professor request.ProfessorRequest) *rest_err.RestErr
	Update(professor request.AtualizarProfessorRequest) *rest_err.RestErr
	Delete(professorId uint) *rest_err.RestErr
	FindById(professorId uint) (response.ProfessorResponse, *rest_err.RestErr)
	FindAll() ([]response.ProfessorResponse, *rest_err.RestErr)
}
