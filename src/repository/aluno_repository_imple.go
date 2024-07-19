package repository

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/models"

	"gorm.io/gorm"
)

type AlunoRepositoryImple struct {
	Db *gorm.DB
}

func NewAlunoRepositoryImple(Db *gorm.DB) AlunoRepository {
	return &AlunoRepositoryImple{Db: Db}
}

func (a *AlunoRepositoryImple) Save(aluno models.Aluno) *rest_err.RestErr {
	if err := a.Db.Create(&aluno).Error; err != nil {
		return rest_err.NewInternalServerError("Erro ao salvar aluno")
	}
	return nil
}

func (a *AlunoRepositoryImple) Update(aluno models.Aluno) *rest_err.RestErr {
	if err := a.Db.Model(&aluno).Updates(aluno).Error; err != nil {
		return rest_err.NewInternalServerError("Erro ao atualizar aluno")
	}
	return nil
}

func (a *AlunoRepositoryImple) Delete(alunoId uint) *rest_err.RestErr {
	var aluno models.Aluno
	if err := a.Db.Preload("Turmas").Where("id = ?", alunoId).First(&aluno).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return rest_err.NewNotFoundError("Aluno não encontrado")
		}
		return rest_err.NewInternalServerError("Erro ao buscar aluno")
	}

	if err := a.Db.Model(&aluno).Association("Turmas").Clear(); err != nil {
		return rest_err.NewInternalServerError("Erro ao desvincular turmas")
	}

	if err := a.Db.Where("aluno_id = ?", alunoId).Delete(&models.Nota{}).Error; err != nil {
		return rest_err.NewInternalServerError("Erro ao deletar notas do aluno")
	}

	if err := a.Db.Delete(&aluno).Error; err != nil {
		return rest_err.NewInternalServerError("Erro ao deletar aluno")
	}
	return nil
}

func (a *AlunoRepositoryImple) FindById(alunoId uint) (models.Aluno, *rest_err.RestErr) {
	var aluno models.Aluno
	if err := a.Db.Where("id = ?", alunoId).First(&aluno).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return aluno, rest_err.NewNotFoundError("Aluno não encontrado")
		}
		return aluno, rest_err.NewInternalServerError("Erro ao buscar aluno")
	}
	return aluno, nil
}

func (a *AlunoRepositoryImple) FindAll() ([]models.Aluno, *rest_err.RestErr) {
	var alunos []models.Aluno
	if err := a.Db.Find(&alunos).Error; err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao buscar todos os alunos")
	}
	return alunos, nil
}
