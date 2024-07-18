package repository

import "controle-notas/src/models"

type AlunoRepository interface {
	Save(aluno models.Aluno)
	Update(aluno models.Aluno)
	Delete(alunoId uint)
	FindById(alunoId uint) (aluno models.Aluno, err error)
	FindAll() []models.Aluno
}
