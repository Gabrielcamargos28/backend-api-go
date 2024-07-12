package professor

import (
	"controle-notas/src/data/professor/request"
	"controle-notas/src/data/professor/response"
	"controle-notas/src/models"
	"controle-notas/src/repository"

	"github.com/go-playground/validator/v10"
)

type ProfessorServiceImple struct {
	ProfessorRepository repository.ProfessorRepository
	validate            *validator.Validate
}

func NewProfessorServiceImple(professorRepository repository.ProfessorRepository, validate *validator.Validate) ProfessorService {
	return &ProfessorServiceImple{
		ProfessorRepository: professorRepository,
		validate:            validate,
	}
}

func (p *ProfessorServiceImple) Create(professor request.ProfessorRequest) {

	professorModel := models.Professor{
		Nome:  professor.Nome,
		Email: professor.Email,
		CPF:   professor.CPF,
	}
	p.ProfessorRepository.Save(professorModel)
}

func (p *ProfessorServiceImple) Delete(professorId int) {
	p.ProfessorRepository.Delete(professorId)
}

func (p *ProfessorServiceImple) FindAll() []response.ProfessorResponse {
	result := p.ProfessorRepository.FindAll()

	var produtos []response.ProfessorResponse
	for _, value := range result {
		produto := response.ProfessorResponse{
			Id:   value.Id,
			Nome: value.Nome,
		}
		produtos = append(produtos, produto)
	}
	return produtos
}

func (p *ProfessorServiceImple) FindById(professorId int) response.ProfessorResponse {

	professorData, err := p.ProfessorRepository.FindById(professorId)
	error.Error(err)

	professorResponse := response.ProfessorResponse{
		Id:   professorData.Id,
		Nome: professorData.Nome,
	}
	return professorResponse
}

func (p *ProfessorServiceImple) Update(professor request.AtualizaProfessorRequest) {
	professorData, err := p.ProfessorRepository.FindById((professor.Id))
	error.Error(err)
	professorData.Nome = professor.Nome
	professorData.Email = professor.Email
	professorData.CPF = professor.CPF
	p.ProfessorRepository.Update(professorData)
}
