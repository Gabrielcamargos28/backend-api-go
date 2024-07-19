package atividade

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data"
	"controle-notas/src/models"
	"controle-notas/src/repository"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AtividadeServiceImple struct {
	AtividadeRepository repository.AtividadeRepository
	TurmaRepository     repository.TurmaRepository
	validate            *validator.Validate
}

func NewAtividadeServiceImple(atividadeRepository repository.AtividadeRepository, turmaRepository repository.TurmaRepository, validate *validator.Validate) AtividadeService {
	return &AtividadeServiceImple{
		AtividadeRepository: atividadeRepository,
		TurmaRepository:     turmaRepository,
		validate:            validate,
	}
}

func (a *AtividadeServiceImple) Create(atividade data.AtividadeRequest) *rest_err.RestErr {
	atividadeModel := models.Atividade{
		Nome:    atividade.Nome,
		Valor:   atividade.Valor,
		Data:    atividade.Data,
		TurmaId: atividade.TurmaId,
	}

	if err := a.AtividadeRepository.Save(atividadeModel); err != nil {
		return rest_err.NewInternalServerError("Erro ao salvar a atividade")
	}
	return nil
}

func (a *AtividadeServiceImple) Update(atividade data.AtualizarAtividadeRequest) *rest_err.RestErr {
	atividadeModel, err := a.AtividadeRepository.FindById(atividade.Id)
	if err != nil {
		log.Printf("Erro ao buscar atividade por ID %d: %v", atividade.Id, err)
		return rest_err.NewInternalServerError("Erro ao buscar atividade para atualização")
	}

	atividadeModel.Nome = atividade.Nome
	atividadeModel.Valor = atividade.Valor
	atividadeModel.Data = atividade.Data

	if err := a.AtividadeRepository.Update(atividadeModel); err != nil {
		return rest_err.NewInternalServerError("Erro ao atualizar a atividade")
	}
	return nil
}

func (a *AtividadeServiceImple) Delete(atividadeId uint) *rest_err.RestErr {
	if err := a.AtividadeRepository.Delete(atividadeId); err != nil {
		return rest_err.NewInternalServerError("Erro ao deletar a atividade")
	}
	return nil
}

func (a *AtividadeServiceImple) FindById(atividadeId uint) (data.AtividadeResponse, *rest_err.RestErr) {
	atividadeModel, err := a.AtividadeRepository.FindById(atividadeId)
	if err != nil {
		log.Printf("Erro ao buscar atividade por ID %d: %v", atividadeId, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return data.AtividadeResponse{}, rest_err.NewNotFoundError("Atividade não encontrada")
		}
		return data.AtividadeResponse{}, rest_err.NewInternalServerError("Erro ao buscar atividade por ID")
	}

	turmaModel, err := a.TurmaRepository.FindById(atividadeModel.TurmaId)
	if err != nil {
		log.Printf("Erro ao buscar turma por ID %d: %v", atividadeModel.TurmaId, err)
		return data.AtividadeResponse{}, rest_err.NewInternalServerError("Erro ao buscar turma associada à atividade")
	}

	atividadeResponse := data.AtividadeResponse{
		Id:    atividadeModel.Id,
		Nome:  atividadeModel.Nome,
		Valor: atividadeModel.Valor,
		Data:  atividadeModel.Data,
		Turma: data.TurmaResponse{
			Id:        turmaModel.Id,
			Nome:      turmaModel.Nome,
			Semestre:  turmaModel.Semestre,
			Ano:       turmaModel.Ano,
			Professor: turmaModel.Professor.Nome,
		},
	}

	alunosNotas, err := a.AtividadeRepository.FindAlunosNotas(atividadeId)
	if err != nil {
		log.Printf("Erro ao buscar alunos e notas para atividade ID %d: %v", atividadeId, err)
		return data.AtividadeResponse{}, rest_err.NewInternalServerError("Erro ao buscar alunos e notas associados à atividade")
	}

	for _, alunoNota := range alunosNotas {
		atividadeResponse.Notas = append(atividadeResponse.Notas, data.AlunoNota{
			AlunoId: alunoNota.AlunoID,
			Nota:    alunoNota.Nota,
		})
	}

	return atividadeResponse, nil
}

func (a *AtividadeServiceImple) FindAll() ([]data.AtividadeResponse, *rest_err.RestErr) {
	atividadesModel, err := a.AtividadeRepository.FindAll()
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao buscar todas as atividades")
	}

	var atividadesResponse []data.AtividadeResponse
	for _, atividade := range atividadesModel {

		turmaModel, err := a.TurmaRepository.FindById(atividade.TurmaId)
		if err != nil {
			log.Printf("Erro ao buscar turma por ID %d: %v", atividade.TurmaId, err)
			return nil, rest_err.NewInternalServerError("Erro ao buscar turma associada à atividade")
		}

		atividadeResponse := data.AtividadeResponse{
			Id:    atividade.Id,
			Nome:  atividade.Nome,
			Valor: atividade.Valor,
			Data:  atividade.Data,
			Turma: data.TurmaResponse{
				Id:        turmaModel.Id,
				Nome:      turmaModel.Nome,
				Semestre:  turmaModel.Semestre,
				Ano:       turmaModel.Ano,
				Professor: turmaModel.Professor.Nome,
			},
		}

		atividadesResponse = append(atividadesResponse, atividadeResponse)
	}

	return atividadesResponse, nil
}
