package nota

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data"
)

type NotaService interface {
	Create(nota data.NotaRequest) *rest_err.RestErr
	Update(nota data.AtualizarNota) *rest_err.RestErr
	Delete(notaId uint) *rest_err.RestErr
	FindById(notaId uint) (data.AlunoNota, *rest_err.RestErr)
	FindAll() ([]data.AlunoNota, *rest_err.RestErr)
}
