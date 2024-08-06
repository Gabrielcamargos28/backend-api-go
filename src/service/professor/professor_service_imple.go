package professor

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data"
	"controle-notas/src/models"
	"controle-notas/src/repository"
	"strings"

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

func (p *ProfessorServiceImple) Create(professor data.ProfessorRequest) *rest_err.RestErr {
	professorModel := models.Professor{
		Nome:  professor.Nome,
		Email: professor.Email,
		CPF:   professor.CPF,
	}

	err := p.ProfessorRepository.Save(professorModel)
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao salvar o professor")
	}
	return nil
}

func (p *ProfessorServiceImple) Delete(professorId uint) *rest_err.RestErr {
	err := p.ProfessorRepository.Delete(professorId)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return rest_err.NewBadRequestError("Não é possível deletar o professor pois ele está associado a uma ou mais turmas")
		}
		return rest_err.NewInternalServerError("Erro ao deletar o professor")
	}
	return nil
}

func (p *ProfessorServiceImple) FindAll() ([]data.ProfessorResponse, *rest_err.RestErr) {
	result, err := p.ProfessorRepository.FindAll()
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao buscar todos os professores")
	}

	var professors []data.ProfessorResponse
	for _, value := range result {
		professor := data.ProfessorResponse{
			Id:   value.Id,
			Nome: value.Nome,
			Email: value.Email,
		}
		professors = append(professors, professor)
	}
	return professors, nil
}

func (p *ProfessorServiceImple) FindById(professorId uint) (data.ProfessorResponse, *rest_err.RestErr) {
	professorData, err := p.ProfessorRepository.FindById(professorId)
	if err != nil {
		log.Printf("Erro ao buscar professor por ID %d: %v", professorId, err)
		return data.ProfessorResponse{}, rest_err.NewInternalServerError("Erro ao buscar professor por ID")
	}
	professorResponse := data.ProfessorResponse{
		Id:   professorData.Id,
		Nome: professorData.Nome,
	}
	return professorResponse, nil
}

func (p *ProfessorServiceImple) Update(professor data.AtualizarProfessorRequest) *rest_err.RestErr {
	professorData, err := p.ProfessorRepository.FindById(professor.Id)
	if err != nil {
		log.Printf("Erro ao atualizar professor com ID %d: %v", professor.Id, err)
		return rest_err.NewInternalServerError("Erro ao buscar professor para atualização")
	}

	professorData.Nome = professor.Nome
	professorData.Email = professor.Email
	professorData.CPF = professor.CPF

	err = p.ProfessorRepository.Update(professorData)
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao atualizar o professor")
	}
	return nil
}
