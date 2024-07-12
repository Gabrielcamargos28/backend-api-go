package professor

import (
	"controle-notas/src/data/professor/request"
	"controle-notas/src/data/professor/response"
)

type ProfessorService interface {
	Create(professor request.ProfessorRequest)
	Update(professor request.AtualizaProfessorRequest)
	Delete(professorId int)
	FindById(professorId int) response.ProfessorResponse
	FindAll() []response.ProfessorResponse
}