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
	t.Db.Preload("Professor").Find(&turmas)
	return turmas
}

func (t *TurmaRepositoryImple) FindById(turmaId uint) (models.Turma, error) {
	var turma models.Turma
	err := t.Db.Preload("Professor").First(&turma, turmaId).Error
	return turma, err
}

func (t *TurmaRepositoryImple) Save(turma models.Turma) {
	t.Db.Create(&turma)
}

func (t *TurmaRepositoryImple) Update(turma models.Turma) error {
	resultado := t.Db.Model(&turma).Updates(models.Turma{
		Nome:        turma.Nome,
		Semestre:    turma.Semestre,
		Ano:         turma.Ano,
		ProfessorId: turma.ProfessorId,
	})
	if resultado.Error != nil {
		return resultado.Error
	}
	return nil
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
	assoc := r.Db.Model(&turma).Association("Alunos")

	if err := assoc.Delete(models.Aluno{Id: alunoId}); err != nil {
		return err
	}

	return nil
}
