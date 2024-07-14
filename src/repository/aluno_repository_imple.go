package repository

import (
	"controle-notas/src/data/aluno/request"
	"controle-notas/src/models"

	"gorm.io/gorm"
)

type AlunoRepositoryImple struct {
	Db *gorm.DB
}

func NewAlunoRepositoryImple(Db *gorm.DB) AlunoRepository {
	return &AlunoRepositoryImple{Db: Db}
}

func (a *AlunoRepositoryImple) Delete(alunoId uint) {
	var aluno models.Professor
	a.Db.Where("id = ?", alunoId).Delete(&aluno)
}

func (a *AlunoRepositoryImple) FindAll() []models.Aluno {
	var alunos []models.Aluno
	a.Db.Find(&alunos)
	return alunos
}

func (a *AlunoRepositoryImple) FindById(alunoId uint) (models.Aluno, error) {
	var aluno models.Aluno
	err := a.Db.First(&aluno, alunoId).Error
	return aluno, err
}

func (a *AlunoRepositoryImple) Save(aluno models.Aluno) {
	a.Db.Create(&aluno)
}

func (a *AlunoRepositoryImple) Update(aluno models.Aluno) {
	var updateAluno = request.AtualizarAlunoRequest{
		Id:        aluno.Id,
		Nome:      aluno.Nome,
		Matricula: aluno.Matricula,
	}
	a.Db.Model(&aluno).Updates(updateAluno)
}
