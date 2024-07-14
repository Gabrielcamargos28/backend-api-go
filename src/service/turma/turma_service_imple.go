package turma

import (
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

func (t *TurmaServiceImple) Create(turma request.TurmaRequest) {
	turmaModel := models.Turma{
		Nome:        turma.Nome,
		Semestre:    turma.Semestre,
		Ano:         turma.Ano,
		ProfessorId: turma.ProfessorId,
	}
	t.TurmaRepository.Save(turmaModel)
}

func (t *TurmaServiceImple) Delete(turmaId uint) {
	t.TurmaRepository.Delete(turmaId)
}

func (t *TurmaServiceImple) FindAll() []response.TurmaResponse {
	result := t.TurmaRepository.FindAll()

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
	return turmas
}

func (t *TurmaServiceImple) FindById(turmaId uint) response.TurmaResponse {
	turmaData, err := t.TurmaRepository.FindById(turmaId)
	if err != nil {
		log.Printf("Erro ao buscar turma por ID %d: %v", turmaId, err)
		return response.TurmaResponse{}
	}

	turmaResponse := response.TurmaResponse{
		Id:        turmaData.Id,
		Nome:      turmaData.Nome,
		Semestre:  turmaData.Semestre,
		Ano:       turmaData.Ano,
		Professor: turmaData.Professor.Nome,
	}
	return turmaResponse
}

func (t *TurmaServiceImple) Update(turma request.AtualizaTurmaRequest) {
	turmaData, err := t.TurmaRepository.FindById(turma.Id)
	if err != nil {
		log.Printf("Erro ao atualizar: %v", err)
		return
	}
	turmaData.Nome = turma.Nome
	turmaData.Semestre = turma.Semestre
	turmaData.Ano = turma.Ano
	turmaData.ProfessorId = turma.ProfessorId
	t.TurmaRepository.Update(turmaData)
}

func (t *TurmaServiceImple) AdicionarAlunos(request request.AdicioanrAlunosTurma) {
	err := t.validate.Struct(request)

	if err != nil {
		log.Printf("Erro ao validar requisição: %v", err)
		return
	}

	turma, err := t.TurmaRepository.FindById(request.TurmaId)
	if err != nil {
		log.Printf("Erro ao buscar a turma: %v", err)
		return
	}
	for _, alunoId := range request.AlunosId {
		aluno, err := t.AlunoRepository.FindById(alunoId)
		if err != nil {
			log.Printf("Erro ao buscar aluno: %v", err)
			return
		}
		turma.Alunos = append(turma.Alunos, aluno)
	}
	t.TurmaRepository.Update(turma)
	if err != nil {
		log.Printf("Erro ao atualizar turma: %v", err)
	}
}
