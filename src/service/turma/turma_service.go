package turma

import (
	"controle-notas/src/data/turma/request"
	"controle-notas/src/data/turma/response"
)

type TurmaService interface {
	Create(turma request.TurmaRequest)
	Update(turma request.AtualizaTurmaRequest)
	Delete(turmaId uint) (response.TurmaResponse, error)
	FindById(turmaId uint) (response.TurmaResponse, error)
	FindAll() []response.TurmaResponse
	AdicionarAlunos(turma request.AdicioanrAlunosTurma)
	RemoveAlunoTurma(alunoId uint, turmaId uint) error
}
