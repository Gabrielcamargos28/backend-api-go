package repository

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/models"

	"gorm.io/gorm"
)

type TurmaRepositoryImple struct {
	Db *gorm.DB
}

func NewTurmaRepositoryImple(Db *gorm.DB) TurmaRepository {
	return &TurmaRepositoryImple{Db: Db}
}

func (t *TurmaRepositoryImple) Save(turma models.Turma) *rest_err.RestErr {
	if result := t.Db.Create(&turma); result.Error != nil {
		return rest_err.NewInternalServerError("Erro ao salvar turma")
	}
	return nil
}

func (t *TurmaRepositoryImple) Update(turma models.Turma) *rest_err.RestErr {
	if result := t.Db.Model(&turma).Updates(models.Turma{
		Nome:        turma.Nome,
		Semestre:    turma.Semestre,
		Ano:         turma.Ano,
		ProfessorId: turma.ProfessorId,
	}); result.Error != nil {
		return rest_err.NewInternalServerError("Erro ao atualizar turma")
	}
	return nil
}

func (t *TurmaRepositoryImple) Delete(turmaId uint) *rest_err.RestErr {
	var turma models.Turma
	if result := t.Db.Where("id = ?", turmaId).Delete(&turma); result.Error != nil {
		return rest_err.NewInternalServerError("Erro ao deletar turma")
	}
	return nil
}

func (t *TurmaRepositoryImple) FindById(turmaId uint) (models.Turma, *rest_err.RestErr) {
	var turma models.Turma
	result := t.Db.Preload("Alunos").First(&turma, turmaId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return turma, rest_err.NewNotFoundError("Turma não encontrada")
		}
		return turma, rest_err.NewInternalServerError("Erro ao buscar turma")
	}
	return turma, nil
}

func (t *TurmaRepositoryImple) FindAll() ([]models.Turma, *rest_err.RestErr) {
	var turmas []models.Turma
	result := t.Db.Preload("Professor").Find(&turmas)
	if result.Error != nil {
		return nil, rest_err.NewInternalServerError("Erro ao buscar turmas")
	}
	return turmas, nil
}

func (t *TurmaRepositoryImple) RemoveAlunoTurma(turmaId uint, alunoId uint) *rest_err.RestErr {
	var turma models.Turma
	if result := t.Db.Preload("Alunos").First(&turma, turmaId); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return rest_err.NewNotFoundError("Turma não encontrada")
		}
		return rest_err.NewInternalServerError("Erro ao buscar turma")
	}

	var aluno models.Aluno
	if result := t.Db.First(&aluno, alunoId); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return rest_err.NewNotFoundError("Aluno não encontrado")
		}
		return rest_err.NewInternalServerError("Erro ao buscar aluno")
	}

	if err := t.Db.Model(&turma).Association("Alunos").Delete(&aluno); err != nil {
		return rest_err.NewInternalServerError("Erro ao remover aluno da turma")
	}

	return nil
}
