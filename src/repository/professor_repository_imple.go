package repository

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data/professor/request"
	"controle-notas/src/models"

	"gorm.io/gorm"
)

type ProfessorRepositoryImple struct {
	Db *gorm.DB
}

func NewProfessorRepositoryImple(Db *gorm.DB) ProfessorRepository {
	return &ProfessorRepositoryImple{Db: Db}
}

func (p *ProfessorRepositoryImple) Delete(professorId uint) *rest_err.RestErr {
	var professor models.Professor
	if result := p.Db.Where("id = ?", professorId).Delete(&professor); result.Error != nil {
		return rest_err.NewInternalServerError("Erro ao deletar professor")
	}
	return nil
}

func (p *ProfessorRepositoryImple) FindAll() ([]models.Professor, *rest_err.RestErr) {
	var professores []models.Professor
	resultado := p.Db.Find(&professores)
	if resultado.Error != nil {
		return nil, rest_err.NewInternalServerError("Erro ao buscar professores")
	}
	return professores, nil
}

func (p *ProfessorRepositoryImple) FindById(professorId uint) (models.Professor, *rest_err.RestErr) {
	var professor models.Professor
	resultado := p.Db.First(&professor, professorId)
	if resultado.Error != nil {
		if resultado.Error == gorm.ErrRecordNotFound {
			return professor, rest_err.NewInternalServerError("Professor não encontrado")
		}
		return professor, rest_err.NewInternalServerError("Erro ao buscar professor")
	}
	return professor, nil
}

func (p *ProfessorRepositoryImple) Save(professor models.Professor) *rest_err.RestErr {
	if result := p.Db.Create(&professor); result.Error != nil {
		return rest_err.NewInternalServerError("Erro ao salvar professor")
	}
	return nil
}

func (p *ProfessorRepositoryImple) Update(professor models.Professor) *rest_err.RestErr {
	var updateProfessor = request.AtualizarProfessorRequest{
		Id:    professor.Id,
		Nome:  professor.Nome,
		Email: professor.Email,
	}
	if result := p.Db.Model(&professor).Updates(updateProfessor); result.Error != nil {
		return rest_err.NewInternalServerError("Erro ao atualizar professor")
	}
	return nil
}
