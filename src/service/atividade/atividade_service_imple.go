package atividade

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data"
	"controle-notas/src/models"
	"controle-notas/src/repository"
	"log"

	"github.com/go-playground/validator/v10"
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
		return data.AtividadeResponse{}, err
	}

	turmaResponse := data.TurmaResponse{
		Id:        atividadeModel.Turma.Id,
		Nome:      atividadeModel.Turma.Nome,
		Semestre:  atividadeModel.Turma.Semestre,
		Ano:       atividadeModel.Turma.Ano,
		Professor: atividadeModel.Turma.Professor.Nome,
	}

	var notasResponse []data.AlunoNota
	for _, nota := range atividadeModel.Notas {
		notaResponse := data.AlunoNota{
			AlunoId:       nota.AlunoId,
			AlunoNome:     nota.Aluno.Nome,
			Nota:          nota.Valor,
			TurmaId:       atividadeModel.Turma.Id,
			TurmaNome:     atividadeModel.Turma.Nome,
			AtividadeId:   atividadeModel.Id,
			AtividadeNome: atividadeModel.Nome,
		}
		notasResponse = append(notasResponse, notaResponse)
	}

	atividadeResponse := data.AtividadeResponse{
		Id:    atividadeModel.Id,
		Nome:  atividadeModel.Nome,
		Valor: atividadeModel.Valor,
		Data:  atividadeModel.Data,
		Turma: turmaResponse,
		Notas: notasResponse,
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
		turmaResponse := data.TurmaResponse{
			Id:        atividade.Turma.Id,
			Nome:      atividade.Turma.Nome,
			Semestre:  atividade.Turma.Semestre,
			Ano:       atividade.Turma.Ano,
			Professor: atividade.Turma.Professor.Nome,
		}

		var notasResponse []data.AlunoNota
		for _, nota := range atividade.Notas {
			notaResponse := data.AlunoNota{
				AlunoId:       nota.AlunoId,
				AlunoNome:     nota.Aluno.Nome,
				Nota:          nota.Valor,
				TurmaId:       atividade.Turma.Id,
				TurmaNome:     atividade.Turma.Nome,
				AtividadeId:   atividade.Id,
				AtividadeNome: atividade.Nome,
			}
			notasResponse = append(notasResponse, notaResponse)
		}

		atividadeResponse := data.AtividadeResponse{
			Id:    atividade.Id,
			Nome:  atividade.Nome,
			Valor: atividade.Valor,
			Data:  atividade.Data,
			Turma: turmaResponse,
			Notas: notasResponse,
		}

		atividadesResponse = append(atividadesResponse, atividadeResponse)
	}

	return atividadesResponse, nil
}
