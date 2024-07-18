package nota

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data"
	"controle-notas/src/models"
	"controle-notas/src/repository"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
)

type NotaServiceImple struct {
	NotaRepository repository.NotaRepository
	validate       *validator.Validate
}

func NewNotaServiceImple(notaRepository repository.NotaRepository, validate *validator.Validate) NotaService {
	return &NotaServiceImple{
		NotaRepository: notaRepository,
		validate:       validate,
	}
}

func (n *NotaServiceImple) Create(nota data.NotaRequest) *rest_err.RestErr {
	notaModel := models.Nota{
		AlunoId:     nota.AlunoId,
		AtividadeId: nota.AtividadeId,
		Valor:       nota.Valor,
	}
	err := n.NotaRepository.Save(notaModel)
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao salvar o nota")
	}
	return nil

}

func (n *NotaServiceImple) Delete(notaId uint) *rest_err.RestErr {
	err := n.NotaRepository.Delete(notaId)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return rest_err.NewBadRequestError("Não é possível deletar o professor pois ele está associado a uma ou mais turmas")
		}
		return rest_err.NewInternalServerError("Erro ao deletar o professor")
	}
	return nil
}

func (n *NotaServiceImple) FindAll() ([]data.AlunoNota, *rest_err.RestErr) {
	result, err := n.NotaRepository.FindAll()
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao buscar todas as notas")
	}

	var notas []data.AlunoNota
	for _, value := range result {
		nota := data.AlunoNota{
			AlunoId:       value.AlunoId,
			AlunoNome:     value.Aluno.Nome,
			Nota:          value.Valor,
			TurmaId:       value.Atividade.TurmaId,
			TurmaNome:     value.Atividade.Turma.Nome,
			AtividadeId:   value.AtividadeId,
			AtividadeNome: value.Atividade.Nome,
		}
		notas = append(notas, nota)
	}
	return notas, nil
}
func (n *NotaServiceImple) FindById(notaId uint) (data.AlunoNota, *rest_err.RestErr) {
	notaData, err := n.NotaRepository.FindById(notaId)
	if err != nil {
		log.Printf("Erro ao buscar nota por ID %d: %v", notaId, err)
		return data.AlunoNota{}, rest_err.NewInternalServerError("Erro ao buscar nota por ID")
	}

	notaResponse := data.AlunoNota{
		AlunoId:   notaData.AlunoId,
		AlunoNome: notaData.Aluno.Nome,
		Nota:      notaData.Valor,
		TurmaId:   notaData.Atividade.TurmaId,
		TurmaNome: notaData.Atividade.Turma.Nome,
	}
	return notaResponse, nil
}

func (n *NotaServiceImple) Update(nota data.AtualizarNota) *rest_err.RestErr {
	if err := n.validate.Struct(nota); err != nil {
		return rest_err.NewBadRequestError("Dados inválidos")
	}

	notaModel := models.Nota{
		Id:    nota.Id,
		Valor: nota.Valor,
	}

	err := n.NotaRepository.Update(notaModel)
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao atualizar a nota")
	}
	return nil
}
