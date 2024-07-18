package atividade

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data"
	"controle-notas/src/models"
	"controle-notas/src/repository"
	"log"
)

type AtividadeServiceImple struct {
	AtividadeRepository repository.AtividadeRepository
}

func NewAtividadeServiceImple(repo repository.AtividadeRepository) AtividadeService {
	return &AtividadeServiceImple{AtividadeRepository: repo}
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
	atividadeModel.TurmaId = atividade.TurmaId

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
		return data.AtividadeResponse{}, rest_err.NewInternalServerError("Erro ao buscar atividade por ID")
	}

	atividadeResponse := data.AtividadeResponse{
		Id:    atividadeModel.Id,
		Nome:  atividadeModel.Nome,
		Valor: atividadeModel.Valor,
		Data:  atividadeModel.Data,
		Turma: data.TurmaResponse{
			Id:   atividadeModel.Turma.Id,
			Nome: atividadeModel.Turma.Nome,
		},
		AlunosNotas: []data.AlunoNota{},
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
		atividadeResponse := data.AtividadeResponse{
			Id:    atividade.Id,
			Nome:  atividade.Nome,
			Valor: atividade.Valor,
			Data:  atividade.Data,
			Turma: data.TurmaResponse{
				Id:   atividade.Turma.Id,
				Nome: atividade.Turma.Nome,
			},
			AlunosNotas: []data.AlunoNota{},
		}
		atividadesResponse = append(atividadesResponse, atividadeResponse)
	}

	return atividadesResponse, nil
}
