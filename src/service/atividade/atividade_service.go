package atividade

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data"
)

type AtividadeService interface {
	Create(atividade data.AtividadeRequest) *rest_err.RestErr
	Update(atividade data.AtualizarAtividadeRequest) *rest_err.RestErr
	Delete(atividadeId uint) *rest_err.RestErr
	FindById(atividadeId uint) (data.AtividadeResponse, *rest_err.RestErr)
	FindAll() ([]data.AtividadeResponse, *rest_err.RestErr)
}
