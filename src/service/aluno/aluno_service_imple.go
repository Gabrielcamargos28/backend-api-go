package aluno

import (
	"controle-notas/src/data/aluno/request"
	"controle-notas/src/data/aluno/response"
	"controle-notas/src/models"
	"controle-notas/src/repository"
	"log"

	"github.com/go-playground/validator/v10"
)

type AlunoServiceImple struct {
	AlunoRepository repository.AlunoRepository
	validate        *validator.Validate
}

func NewAlunoServiceImple(alunoRepository repository.AlunoRepository, validate *validator.Validate) AlunoService {
	return &AlunoServiceImple{
		AlunoRepository: alunoRepository,
		validate:        validate,
	}
}

func (a *AlunoServiceImple) Create(aluno request.AlunoRequest) {

	if err := a.validate.Struct(aluno); err != nil {
		log.Printf("Erro de validação: %v", err)
		return
	}

	alunoModel := models.Aluno{
		Nome:      aluno.Nome,
		Matricula: aluno.Matricula,
	}
	a.AlunoRepository.Save(alunoModel)
}

func (a *AlunoServiceImple) Delete(alunoId uint) {
	a.AlunoRepository.Delete(alunoId)
}

func (a *AlunoServiceImple) FindAll() []response.AlunoResponse {
	result := a.AlunoRepository.FindAll()
	var alunos []response.AlunoResponse
	for _, value := range result {
		aluno := response.AlunoResponse{
			Id:   value.Id,
			Nome: value.Nome,
		}
		alunos = append(alunos, aluno)
	}
	return alunos
}

func (a *AlunoServiceImple) FindById(alunoId uint) response.AlunoResponse {
	alunoData, err := a.AlunoRepository.FindById(alunoId)
	if err != nil {
		log.Printf("Erro ao buscar aluno pelo ID %d: %v", alunoId, err)
		return response.AlunoResponse{}
	}
	alunoResponse := response.AlunoResponse{
		Id:   alunoData.Id,
		Nome: alunoData.Nome,
	}
	return alunoResponse
}

func (a *AlunoServiceImple) Update(aluno request.AtualizarAlunoRequest) {

	if err := a.validate.Struct(aluno); err != nil {
		log.Printf("Erro de validação: %v", err)
		return
	}

	alunoData, err := a.AlunoRepository.FindById(aluno.Id)
	if err != nil {
		log.Printf("Erro ao atualizar aluno com ID %d: %v", aluno.Id, err)
		return
	}
	alunoData.Nome = aluno.Nome
	alunoData.Matricula = aluno.Matricula
	a.AlunoRepository.Update(alunoData)
}
