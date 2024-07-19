package aluno

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data"
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

func (a *AlunoServiceImple) Create(aluno data.AlunoRequest) *rest_err.RestErr {
	if err := a.validate.Struct(aluno); err != nil {
		log.Printf("Erro de validação: %v", err)
		return rest_err.NewBadRequestError("Erro de validação")
	}

	alunoModel := models.Aluno{
		Nome:      aluno.Nome,
		Matricula: aluno.Matricula,
	}
	return a.AlunoRepository.Save(alunoModel)
}

func (a *AlunoServiceImple) Update(aluno data.AtualizarAlunoRequest) *rest_err.RestErr {
	if err := a.validate.Struct(aluno); err != nil {
		log.Printf("Erro de validação: %v", err)
		return rest_err.NewBadRequestError("Erro de validação")
	}

	alunoData, err := a.AlunoRepository.FindById(aluno.Id)
	if err != nil {
		log.Printf("Erro ao atualizar aluno com ID %d: %v", aluno.Id, err)
		return err
	}
	alunoData.Nome = aluno.Nome
	alunoData.Matricula = aluno.Matricula
	return a.AlunoRepository.Update(alunoData)
}

func (a *AlunoServiceImple) Delete(alunoId uint) *rest_err.RestErr {
	return a.AlunoRepository.Delete(alunoId)
}

func (a *AlunoServiceImple) FindAll() ([]data.AlunoResponse, *rest_err.RestErr) {
	result, err := a.AlunoRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var alunos []data.AlunoResponse
	for _, value := range result {
		aluno := data.AlunoResponse{
			Id:   value.Id,
			Nome: value.Nome,
		}
		alunos = append(alunos, aluno)
	}
	return alunos, nil
}

func (a *AlunoServiceImple) FindById(alunoId uint) (data.AlunoResponse, *rest_err.RestErr) {
	alunoData, err := a.AlunoRepository.FindById(alunoId)
	if err != nil {
		log.Printf("Erro ao buscar aluno pelo ID %d: %v", alunoId, err)
		return data.AlunoResponse{}, rest_err.NewInternalServerError("Erro")
	}
	alunoResponse := data.AlunoResponse{
		Id:   alunoData.Id,
		Nome: alunoData.Nome,
	}
	return alunoResponse, nil
}
