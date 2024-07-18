// src/repository/atividade_repository.go
package repository

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data"
	"controle-notas/src/models"
	"strings"

	"gorm.io/gorm"
)

type AtividadeRepositoryImple struct {
	Db *gorm.DB
}

func NewAtividadeRepositoryImple(Db *gorm.DB) AtividadeRepository {
	return &AtividadeRepositoryImple{Db: Db}
}

func (a *AtividadeRepositoryImple) Delete(atividadeId uint) *rest_err.RestErr {
	var atividade models.Atividade

	result := a.Db.Where("id = ?", atividadeId).First(&atividade)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return rest_err.NewNotFoundError("Atividade não encontrada")
		}
		return rest_err.NewInternalServerError("Erro ao buscar atividade")
	}

	result = a.Db.Delete(&atividade)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "violates foreign key constraint") {
			return rest_err.NewBadRequestError("Não é possível deletar a atividade pois ela está associada a uma ou mais turmas")
		}
		return rest_err.NewInternalServerError("Erro ao deletar atividade")
	}

	return nil
}

func (a *AtividadeRepositoryImple) FindAll() ([]models.Atividade, *rest_err.RestErr) {
	var atividades []models.Atividade
	resultado := a.Db.Find(&atividades)
	if resultado.Error != nil {
		return nil, rest_err.NewInternalServerError("Erro ao buscar atividades")
	}
	return atividades, nil
}

func (a *AtividadeRepositoryImple) FindById(atividadeId uint) (models.Atividade, *rest_err.RestErr) {
	var atividade models.Atividade
	resultado := a.Db.First(&atividade, atividadeId)
	if resultado.Error != nil {
		if resultado.Error == gorm.ErrRecordNotFound {
			return atividade, rest_err.NewInternalServerError("Erro ao encontrar atividade")
		}
		return atividade, rest_err.NewInternalServerError("Erro ao buscar atividade")
	}
	return atividade, nil
}

func (a *AtividadeRepositoryImple) Save(atividade models.Atividade) *rest_err.RestErr {
	if result := a.Db.Create(&atividade); result.Error != nil {
		return rest_err.NewInternalServerError("Erro ao salvar atividade")
	}
	return nil
}

func (a *AtividadeRepositoryImple) Update(request models.Atividade) *rest_err.RestErr {
	var updateAtividade = data.AtualizarAtividadeRequest{
		Id:    request.Id,
		Nome:  request.Nome,
		Valor: request.Valor,
		Data:  request.Data,
	}
	if result := a.Db.Model(&request).Updates(updateAtividade); result.Error != nil {
		return rest_err.NewInternalServerError("Erro ao atualizar atividade")
	}
	return nil
}

func (a *AtividadeRepositoryImple) FindAlunosNotas(atividadeId uint) ([]models.AlunoNota, *rest_err.RestErr) {
	var alunosNotas []models.AlunoNota
	result := a.Db.Where("atividade_id = ?", atividadeId).Find(&alunosNotas)
	if result.Error != nil {
		return nil, rest_err.NewInternalServerError("Erro ao buscar alunos e notas")
	}
	return alunosNotas, nil
}
