package repository

import "controle-notas/src/models"

type TurmaRepository interface {
	Save(turma models.Turma) error
	Update(turma models.Turma) error
	Delete(turmaId uint) error
	FindById(turmaId uint) (models.Turma, error)
	FindAll() ([]models.Turma, error)
	RemoveAlunoTurma(turmaId uint, alunoId uint) error
}
