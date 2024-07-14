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

func (t *TurmaServiceImple) Create(turma request.TurmaRequest) error {
	turmaModel := models.Turma{
		Nome:        turma.Nome,
		Semestre:    turma.Semestre,
		Ano:         turma.Ano,
		ProfessorId: turma.ProfessorId,
	}
	return t.TurmaRepository.Save(turmaModel)
}

func (t *TurmaServiceImple) Delete(turmaId uint) error {
	return t.TurmaRepository.Delete(turmaId)
}

func (t *TurmaServiceImple) FindAll() ([]response.TurmaResponse, error) {
	result, err := t.TurmaRepository.FindAll()
	if err != nil {
		return nil, err
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

func (t *TurmaServiceImple) FindById(turmaId uint) (response.TurmaResponse, error) {
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

func (t *TurmaServiceImple) Update(turma request.AtualizaTurmaRequest) error {
	turmaData, err := t.TurmaRepository.FindById(turma.Id)
	if err != nil {
		log.Printf("Erro ao atualizar: %v", err)
		return err
	}
	turmaData.Nome = turma.Nome
	turmaData.Semestre = turma.Semestre
	turmaData.Ano = turma.Ano
	turmaData.ProfessorId = turma.ProfessorId
	return t.TurmaRepository.Update(turmaData)
}

func (t *TurmaServiceImple) AdicionarAlunos(request request.AdicioanrAlunosTurma) error {
	err := t.validate.Struct(request)
	if err != nil {
		log.Printf("Erro ao validar requisição: %v", err)
		return err
	}

	turma, err := t.TurmaRepository.FindById(request.TurmaId)
	if err != nil {
		log.Printf("Erro ao buscar a turma: %v", err)
		return err
	}

	for _, alunoId := range request.AlunosId {
		aluno, err := t.AlunoRepository.FindById(alunoId)
		if err != nil {
			log.Printf("Erro ao buscar aluno: %v", err)
			return err
		}
		turma.Alunos = append(turma.Alunos, aluno)
	}

	return t.TurmaRepository.Update(turma)
}

func (t *TurmaServiceImple) RemoveAlunoTurma(alunoId uint, turmaId uint) error {
	turma, err := t.TurmaRepository.FindById(turmaId)
	if err != nil {
		return err
	}

	for i, aluno := range turma.Alunos {
		if aluno.Id == alunoId {
			turma.Alunos = append(turma.Alunos[:i], turma.Alunos[i+1:]...)
			break
		}
	}

	return t.TurmaRepository.RemoveAlunoTurma(turmaId, alunoId)
}
