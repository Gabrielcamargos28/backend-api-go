package turma

import (
	"controle-notas/src/data/turma/request"
	"controle-notas/src/data/turma/response"
)

type TurmaService interface {
	Create(turma request.TurmaRequest)
	Update(turma request.AtualizaTurmaRequest)
	Delete(turmaId uint)
	FindById(turmaId uint) response.TurmaResponse
	FindAll() []response.TurmaResponse
	AdicionarAlunos(turma request.AdicioanrAlunosTurma)
}
