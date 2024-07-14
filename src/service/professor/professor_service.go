package professor

import (
	"controle-notas/src/data/professor/request"
	"controle-notas/src/data/professor/response"
)

type ProfessorService interface {
	Create(professor request.ProfessorRequest)
	Update(professor request.AtualizarProfessorRequest)
	Delete(professorId uint)
	FindById(professorId uint) response.ProfessorResponse
	FindAll() []response.ProfessorResponse
}
