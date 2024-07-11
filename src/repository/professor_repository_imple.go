package repository

import (
	"controle-notas/src/models"

	"gorm.io/gorm"
)

type ProfessorRepositoryImple struct {
	Db *gorm.DB
}

func NewProfessorRepositoryImple(Db *gorm.DB) ProfessorRepository {
	return &ProfessorRepositoryImple{Db: Db}
}

func (p *ProfessorRepositoryImple) Delete(professorId int) {
	var professor models.Professor
	p.Db.Where("id = ?", professorId).Delete(&professor)
}

func (p *ProfessorRepositoryImple) FindAll() []models.Professor {
	var professors []models.Professor
	p.Db.Find(&professors)
	return professors
}

func (p *ProfessorRepositoryImple) FindById(professorId int) (models.Professor, error) {
	var professor models.Professor
	err := p.Db.First(&professor, professorId).Error
	return professor, err
}

func (p *ProfessorRepositoryImple) Save(professor models.Professor) {
	p.Db.Create(&professor)
}

func (p *ProfessorRepositoryImple) Update(professor models.Professor) {
	p.Db.Save(&professor)
}
