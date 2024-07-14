package turma

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data/turma/request"
	"controle-notas/src/data/turma/response"
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

func (t *TurmaServiceImple) Create(turma request.TurmaRequest) *rest_err.RestErr {
	turmaModel := models.Turma{
		Nome:        turma.Nome,
		Semestre:    turma.Semestre,
		Ano:         turma.Ano,
		ProfessorId: turma.ProfessorId,
	}
	err := t.TurmaRepository.Save(turmaModel)
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao salvar a turma", nil)
	}
	return nil
}

func (t *TurmaServiceImple) Delete(turmaId uint) *rest_err.RestErr {
	err := t.TurmaRepository.Delete(turmaId)
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao deletar turma", nil)
	}
	return nil
}

func (t *TurmaServiceImple) FindAll() ([]response.TurmaResponse, *rest_err.RestErr) {
	result, err := t.TurmaRepository.FindAll()
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao buscar turmas", nil)
	}

	var turmas []response.TurmaResponse
	for _, value := range result {
		turma := response.TurmaResponse{
			Id:        value.Id,
			Nome:      value.Nome,
			Semestre:  value.Semestre,
			Ano:       value.Ano,
			Professor: value.Professor.Nome,
		}
		turmas = append(turmas, turma)
	}
	return turmas, nil
}

func (t *TurmaServiceImple) FindById(turmaId uint) (response.TurmaResponse, *rest_err.RestErr) {
	turmaData, err := t.TurmaRepository.FindById(turmaId)
	if err != nil {
		log.Printf("Erro ao buscar turma por ID %d: %v", turmaId, err)
		return response.TurmaResponse{}, err
	}

	turmaResponse := response.TurmaResponse{
		Id:        turmaData.Id,
		Nome:      turmaData.Nome,
		Semestre:  turmaData.Semestre,
		Ano:       turmaData.Ano,
		Professor: turmaData.Professor.Nome,
	}
	return turmaResponse, nil
}

func (t *TurmaServiceImple) Update(turma request.AtualizaTurmaRequest) *rest_err.RestErr {
	turmaData, err := t.TurmaRepository.FindById(turma.Id)
	if err != nil {
		log.Printf("Erro ao atualizar: %v", err)
		return rest_err.NewInternalServerError("Erro ao buscar turma por ID", nil)
	}
	turmaData.Nome = turma.Nome
	turmaData.Semestre = turma.Semestre
	turmaData.Ano = turma.Ano
	turmaData.ProfessorId = turma.ProfessorId
	err = t.TurmaRepository.Update(turmaData)
	if err != nil {
		return rest_err.NewInternalServerError("Erroa o atualizar turma", nil)
	}
	return nil
}

func (t *TurmaServiceImple) AdicionarAlunos(request request.AdicioanrAlunosTurma) *rest_err.RestErr {
	err := t.validate.Struct(request)
	if err != nil {
		log.Printf("Erro ao validar requisição: %v", err)
		return rest_err.NewInternalServerError("Erro ao adicionar alunos a turma", nil)
	}

	turma, err := t.TurmaRepository.FindById(request.TurmaId)
	if err != nil {
		log.Printf("Erro ao buscar a turma: %v", err)
		return rest_err.NewInternalServerError("Erro ao buscar turma", nil)
	}

	for _, alunoId := range request.AlunosId {
		aluno, err := t.AlunoRepository.FindById(alunoId)
		if err != nil {
			log.Printf("Erro ao buscar aluno: %v", err)
			return rest_err.NewInternalServerError("Erro ao buscar aluno", nil)
		}
		turma.Alunos = append(turma.Alunos, aluno)
	}

	return t.TurmaRepository.Update(turma)
}

func (t *TurmaServiceImple) RemoveAlunoTurma(alunoId uint, turmaId uint) *rest_err.RestErr {
	turma, err := t.TurmaRepository.FindById(turmaId)
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao buscar turma", nil)
	}

	for i, aluno := range turma.Alunos {
		if aluno.Id == alunoId {
			turma.Alunos = append(turma.Alunos[:i], turma.Alunos[i+1:]...)
			break
		}
	}

	return t.TurmaRepository.RemoveAlunoTurma(turmaId, alunoId)
}
