package repository

import (
	"controle-notas/src/models"

	"gorm.io/gorm"
)

type TurmaRepositoryImple struct {
	Db *gorm.DB
}

func NewTurmaRepositoryImple(Db *gorm.DB) TurmaRepository {
	return &TurmaRepositoryImple{Db: Db}
}

func (t *TurmaRepositoryImple) Save(turma models.Turma) error {
	err := t.Db.Create(&turma).Error
	return err
}

func (t *TurmaRepositoryImple) Update(turma models.Turma) error {
	resultado := t.Db.Model(&turma).Updates(models.Turma{
		Nome:        turma.Nome,
		Semestre:    turma.Semestre,
		Ano:         turma.Ano,
		ProfessorId: turma.ProfessorId,
	})
	return resultado.Error
}

func (t *TurmaRepositoryImple) Delete(turmaId uint) error {
	var turma models.Turma
	resultado := t.Db.Where("id = ?", turmaId).Delete(&turma)
	return resultado.Error
}

func (t *TurmaRepositoryImple) FindById(turmaId uint) (models.Turma, error) {
	var turma models.Turma
	resultado := t.Db.Preload("Professor").First(&turma, turmaId)
	return turma, resultado.Error
}

func (t *TurmaRepositoryImple) FindAll() ([]models.Turma, error) {
	var turmas []models.Turma
	resultado := t.Db.Preload("Professor").Find(&turmas)
	return turmas, resultado.Error
}

func (r *TurmaRepositoryImple) RemoveAlunoTurma(turmaId uint, alunoId uint) error {
	var turma models.Turma
	resultado := r.Db.Preload("Alunos").First(&turma, turmaId)
	if resultado.Error != nil {
		return resultado.Error
	}

	var updatedAlunos []models.Aluno
	for _, aluno := range turma.Alunos {
		if aluno.Id != alunoId {
			updatedAlunos = append(updatedAlunos, aluno)
		}
	}
	turma.Alunos = updatedAlunos

	assoc := r.Db.Model(&turma).Association("Alunos")
	err := assoc.Replace(&turma.Alunos)
	return err
}
