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
	NotaRepository      repository.NotaRepository
	AtividadeRepository repository.AtividadeRepository
	validate            *validator.Validate
}

func NewNotaServiceImple(notaRepository repository.NotaRepository, atividadeRepository repository.AtividadeRepository, validate *validator.Validate) NotaService {
	return &NotaServiceImple{
		NotaRepository:      notaRepository,
		AtividadeRepository: atividadeRepository,
		validate:            validate,
	}
}

func (n *NotaServiceImple) Create(nota data.NotaRequest) *rest_err.RestErr {

	existingNota, err := n.NotaRepository.FindByAlunoAndAtividade(nota.AlunoId, nota.AtividadeId)
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao verificar se a nota já existe")
	}
	if existingNota != nil {
		return rest_err.NewBadRequestError("Nota para o aluno e a atividade já existe")
	}

	notaModel := models.Nota{
		AlunoId:     nota.AlunoId,
		AtividadeId: nota.AtividadeId,
		Valor:       nota.Valor,
	}

	modelAtividade, errAt := n.AtividadeRepository.FindById(notaModel.AtividadeId)

	if errAt != nil {
		return rest_err.NewInternalServerError("Erro ao carregar atividade para busca de nota")
	}

	if notaModel.Valor > modelAtividade.Valor {
		return rest_err.NewBadRequestError("Nota atribuida maior que a nota da atividade")
	}
	err = n.NotaRepository.Save(notaModel)
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao salvar o nota")
	}
	return nil

}

func (n *NotaServiceImple) Delete(notaId uint) *rest_err.RestErr {
	err := n.NotaRepository.Delete(notaId)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return rest_err.NewBadRequestError("Não é possível deletar a nota pois ele está associado a uma ou mais turmas")
		}
		return rest_err.NewInternalServerError("Erro ao deletar a nota")
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
			AlunoId:        value.AlunoId,
			AlunoNome:      value.Aluno.Nome,
			NotaId:         value.Id,
			Nota:           value.Valor,
			TurmaId:        value.Atividade.TurmaId,
			TurmaNome:      value.Atividade.Turma.Nome,
			AtividadeId:    value.AtividadeId,
			AtividadeNome:  value.Atividade.Nome,
			AtividadeValor: value.Atividade.Valor,
		}
		notas = append(notas, nota)
	}
	return notas, nil
}

/*func (n *NotaServiceImple) FindById(notaId uint) (data.AlunoNota, *rest_err.RestErr) {
	notaData, err := n.NotaRepository.FindById(notaId)
	if err != nil {
		log.Printf("Erro ao buscar nota por ID %d: %v", notaId, err)
		return data.AlunoNota{}, rest_err.NewInternalServerError("Erro ao buscar nota por ID")
	}

	notaResponse := data.AlunoNota{
		AlunoId:       notaData.AlunoId,
		AlunoNome:     notaData.Aluno.Nome,
		Nota:          notaData.Valor,
		TurmaId:       notaData.Atividade.Turma.Id,
		TurmaNome:     notaData.Atividade.Turma.Nome,
		AtividadeId:   notaData.AtividadeId,
		AtividadeNome: notaData.Atividade.Nome,
	}
	fmt.Println(notaData)
	return notaResponse, nil
}*/

func (n *NotaServiceImple) FindById(notaId uint) (data.AlunoNota, *rest_err.RestErr) {
	notaData, err := n.NotaRepository.FindById(notaId)
	if err != nil {
		log.Printf("Erro ao buscar nota por ID %d: %v", notaId, err)
		return data.AlunoNota{}, rest_err.NewInternalServerError("Erro ao buscar nota por ID")
	}

	notaResponse := data.AlunoNota{
		AlunoId:        notaData.AlunoId,
		AlunoNome:      notaData.Aluno.Nome,
		NotaId:         notaData.Id,
		Nota:           notaData.Valor,
		TurmaId:        notaData.Atividade.Turma.Id,
		TurmaNome:      notaData.Atividade.Turma.Nome,
		AtividadeId:    notaData.AtividadeId,
		AtividadeNome:  notaData.Atividade.Nome,
		AtividadeValor: notaData.Atividade.Valor,
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
func (n *NotaServiceImple) FindNotasByAlunoId(alunoId uint) ([]data.AlunoNota, *rest_err.RestErr) {
	notas, err := n.NotaRepository.FindNotasByAlunoId(alunoId)
	if err != nil {
		return nil, err
	}

	var notasResponse []data.AlunoNota
	for _, nota := range notas {
		notasResponse = append(notasResponse, data.AlunoNota{
			AlunoId:        nota.AlunoId,
			AlunoNome:      nota.Aluno.Nome,
			NotaId:         nota.Id,
			Nota:           nota.Valor,
			TurmaId:        nota.Atividade.TurmaId,
			TurmaNome:      nota.Atividade.Turma.Nome,
			AtividadeId:    nota.AtividadeId,
			AtividadeNome:  nota.Atividade.Nome,
			AtividadeValor: nota.Atividade.Valor,
		})
	}
	return notasResponse, nil
}
