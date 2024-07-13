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

func (t *TurmaRepositoryImple) Delete(turmaId uint) {
	var turma models.Turma
	t.Db.Where("id = ?", turmaId).Delete(&turma)
}

func (t *TurmaRepositoryImple) FindAll() []models.Turma {
	var turmas []models.Turma
	t.Db.Preload("Professor").Find(&turmas) // Preload to load the related Professor
	return turmas
}

func (t *TurmaRepositoryImple) FindById(turmaId uint) (models.Turma, error) {
	var turma models.Turma
	err := t.Db.Preload("Professor").First(&turma, turmaId).Error // Preload to load the related Professor
	return turma, err
}

func (t *TurmaRepositoryImple) Save(turma models.Turma) {
	t.Db.Create(&turma)
}

func (t *TurmaRepositoryImple) Update(turma models.Turma) {
	t.Db.Model(&turma).Updates(models.Turma{
		Nome:        turma.Nome,
		Semestre:    turma.Semestre,
		Ano:         turma.Ano,
		ProfessorId: turma.ProfessorId,
	})
}
