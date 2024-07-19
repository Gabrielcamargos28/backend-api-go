package turma

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data"
	"controle-notas/src/models"
	"controle-notas/src/repository"
	"log"

	"github.com/go-playground/validator/v10"
)

type TurmaServiceImple struct {
	TurmaRepository repository.TurmaRepository
	AlunoRepository repository.AlunoRepository
	validate        *validator.Validate
}

func NewTurmaServiceImple(turmaRepository repository.TurmaRepository, alunoRepository repository.AlunoRepository, validate *validator.Validate) TurmaService {
	return &TurmaServiceImple{
		TurmaRepository: turmaRepository,
		AlunoRepository: alunoRepository,
		validate:        validate,
	}
}

func (t *TurmaServiceImple) Create(turma data.TurmaRequest) *rest_err.RestErr {
	turmaModel := models.Turma{
		Nome:        turma.Nome,
		Semestre:    turma.Semestre,
		Ano:         turma.Ano,
		ProfessorId: turma.ProfessorId,
	}
	err := t.TurmaRepository.Save(turmaModel)
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao salvar a turma")
	}
	return nil
}

func (t *TurmaServiceImple) Delete(turmaId uint) *rest_err.RestErr {
	err := t.TurmaRepository.Delete(turmaId)
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao deletar turma")
	}
	return nil
}

func (t *TurmaServiceImple) FindAll() ([]data.TurmaResponse, *rest_err.RestErr) {
	result, err := t.TurmaRepository.FindAll()
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao buscar turmas")
	}

	var turmas []data.TurmaResponse
	for _, value := range result {
		turma := data.TurmaResponse{
			Id:          value.Id,
			Nome:        value.Nome,
			Semestre:    value.Semestre,
			Ano:         value.Ano,
			ProfessorId: value.Professor.Id,
			Professor:   value.Professor.Nome,
		}
		turmas = append(turmas, turma)
	}
	return turmas, nil
}

func (t *TurmaServiceImple) FindById(turmaId uint) (data.TurmaAlunosResponse, *rest_err.RestErr) {
	turmaData, err := t.TurmaRepository.FindById(turmaId)
	if err != nil {
		log.Printf("Erro ao buscar turma por ID %d: %v", turmaId, err)
		return data.TurmaAlunosResponse{}, rest_err.NewInternalServerError("Erro ao buscar turma por ID")
	}

	var alunosResponse []data.AlunoResumido
	for _, aluno := range turmaData.Alunos {
		alunosResponse = append(alunosResponse, data.AlunoResumido{
			Id:   aluno.Id,
			Nome: aluno.Nome,
		})
	}

	var somaValores float64
	atividadesTurma, errR := t.TurmaRepository.FindAtividadesByTurmaId(turmaData.Id)
	if errR != nil {
		return data.TurmaAlunosResponse{}, rest_err.NewInternalServerError("Erro ao buscar turma por ID")
	}
	for _, atividade := range atividadesTurma {
		somaValores += atividade.Valor
	}

	var atividadeResponse []data.AtividadeTurmaResponse
	for _, atividade := range turmaData.Atividades {
		atividadeResponse = append(atividadeResponse, data.AtividadeTurmaResponse{
			Id:    atividade.Id,
			Nome:  atividade.Nome,
			Valor: atividade.Valor,
			Data:  atividade.Data,
		})
	}

	turmaResponse := data.TurmaAlunosResponse{
		Id:             turmaData.Id,
		Nome:           turmaData.Nome,
		Semestre:       turmaData.Semestre,
		Ano:            turmaData.Ano,
		ProfessorId:    turmaData.Professor.Id,
		Professor:      turmaData.Professor.Nome,
		Alunos:         alunosResponse,
		Atividades:     atividadeResponse,
		SomaAtividades: somaValores,
	}

	return turmaResponse, nil
}

func (t *TurmaServiceImple) Update(turma data.AtualizaTurmaRequest) *rest_err.RestErr {
	turmaData, err := t.TurmaRepository.FindById(turma.Id)
	if err != nil {
		log.Printf("Erro ao atualizar: %v", err)
		return rest_err.NewInternalServerError("Erro ao buscar turma por ID")
	}
	turmaData.Nome = turma.Nome
	turmaData.Semestre = turma.Semestre
	turmaData.Ano = turma.Ano
	turmaData.ProfessorId = turma.ProfessorId
	err = t.TurmaRepository.Update(turmaData)
	if err != nil {
		return rest_err.NewInternalServerError("Erroa o atualizar turma")
	}
	return nil
}

func (t *TurmaServiceImple) AdicionarAlunos(request data.AdicioanarAlunosTurma) *rest_err.RestErr {
	err := t.validate.Struct(request)
	if err != nil {
		log.Printf("Erro ao validar requisição: %v", err)
		return rest_err.NewInternalServerError("Erro ao validar requisição para adicionar alunos a turma")
	}

	turma, err := t.TurmaRepository.FindById(request.TurmaId)

	for _, alunoId := range request.AlunosId {
		aluno, err := t.AlunoRepository.FindById(alunoId)
		if err != nil {
			log.Printf("Erro ao buscar aluno com ID %d: %v", alunoId, err)
			return rest_err.NewNotFoundError("Aluno não encontrado")
		}
		turma.Alunos = append(turma.Alunos, aluno)
	}

	err = t.TurmaRepository.Update(turma)
	return nil
}

func (t *TurmaServiceImple) RemoveAlunoTurma(alunoId uint, turmaId uint) *rest_err.RestErr {
	err := t.TurmaRepository.RemoveAlunoTurma(turmaId, alunoId)
	if err != nil {
		return err
	}

	return nil
}
func (t *TurmaServiceImple) FindAtividadesByTurmaId(turmaId uint) ([]data.AtividadeTurmaResponse, *rest_err.RestErr) {

	turma, err := t.TurmaRepository.FindById(turmaId)
	if err != nil {
		return nil, err
	}
	var somaValores float64
	atividadesTurma, errR := t.TurmaRepository.FindAtividadesByTurmaId(turma.Id)
	if errR != nil {
		return nil, rest_err.NewInternalServerError("Erro ao buscar turma por ID")
	}
	for _, atividade := range atividadesTurma {
		somaValores += atividade.Valor
	}

	var atividadesResponse []data.AtividadeTurmaResponse
	for _, atividade := range turma.Atividades {
		atividadesResponse = append(atividadesResponse, data.AtividadeTurmaResponse{
			Id:    atividade.Id,
			Nome:  atividade.Nome,
			Valor: atividade.Valor,
			Data:  atividade.Data,
		})
	}

	return atividadesResponse, nil
}
