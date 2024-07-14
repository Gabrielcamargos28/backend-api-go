package professor

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data/professor/request"
	"controle-notas/src/data/professor/response"
	"controle-notas/src/models"
	"controle-notas/src/repository"

	"log"

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

func (p *ProfessorServiceImple) Create(professor request.ProfessorRequest) *rest_err.RestErr {
	professorModel := models.Professor{
		Nome:  professor.Nome,
		Email: professor.Email,
		CPF:   professor.CPF,
	}

	err := p.ProfessorRepository.Save(professorModel)
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao salvar o professor", nil)
	}
	return nil
}

func (p *ProfessorServiceImple) Delete(professorId uint) *rest_err.RestErr {
	err := p.ProfessorRepository.Delete(professorId)
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao deletar o professor", nil)
	}
	return nil
}

func (p *ProfessorServiceImple) FindAll() ([]response.ProfessorResponse, *rest_err.RestErr) {
	result, err := p.ProfessorRepository.FindAll()
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao buscar todos os professores", nil)
	}

	var professors []response.ProfessorResponse
	for _, value := range result {
		professor := response.ProfessorResponse{
			Id:   value.Id,
			Nome: value.Nome,
		}
		professors = append(professors, professor)
	}
	return professors, nil
}

func (p *ProfessorServiceImple) FindById(professorId uint) (response.ProfessorResponse, *rest_err.RestErr) {
	professorData, err := p.ProfessorRepository.FindById(professorId)
	if err != nil {
		log.Printf("Erro ao buscar professor por ID %d: %v", professorId, err)
		return response.ProfessorResponse{}, rest_err.NewInternalServerError("Erro ao buscar professor por ID", nil)
	}

	professorResponse := response.ProfessorResponse{
		Id:   professorData.Id,
		Nome: professorData.Nome,
	}
	return professorResponse, nil
}
func (p *ProfessorServiceImple) Update(professor request.AtualizarProfessorRequest) *rest_err.RestErr {
	professorData, err := p.ProfessorRepository.FindById(professor.Id)
	if err != nil {
		log.Printf("Erro ao atualizar professor com ID %d: %v", professor.Id, err)
		return rest_err.NewInternalServerError("Erro ao buscar professor para atualização", nil)
	}

	professorData.Nome = professor.Nome
	professorData.Email = professor.Email
	professorData.CPF = professor.CPF

	err = p.ProfessorRepository.Update(professorData)
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao atualizar o professor", nil)
	}
	return nil
}
