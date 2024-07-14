package turma

import (
	"controle-notas/src/data/turma/request"
	"controle-notas/src/data/turma/response"
)

type TurmaService interface {
	Create(turma request.TurmaRequest) error
	Update(turma request.AtualizaTurmaRequest) error
	Delete(turmaId uint) error
	FindById(turmaId uint) (response.TurmaResponse, error)
	FindAll() ([]response.TurmaResponse, error)
	AdicionarAlunos(turma request.AdicioanrAlunosTurma) error
	RemoveAlunoTurma(alunoId uint, turmaId uint) error
}
