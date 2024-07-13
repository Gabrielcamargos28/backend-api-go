package repository

import "controle-notas/src/models"

type TurmaRepository interface {
	Save(turma models.Turma)
	Update(turma models.Turma)
	Delete(turmaId uint)
	FindById(turmaId uint) (turma models.Turma, err error)
	FindAll() []models.Turma
}
