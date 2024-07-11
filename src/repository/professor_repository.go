package repository

import "controle-notas/src/models"

type ProfessorRepository interface {
	Save(professor models.Professor)
	Update(professor models.Professor)
	Delete(professorId int)
	FindById(professorId int) (professor models.Professor, err error)
	FindAll() []models.Professor
}
