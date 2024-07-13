package repository

import (
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

func (p *ProfessorRepositoryImple) Delete(professorId uint) {
	var professor models.Professor
	p.Db.Where("id = ?", professorId).Delete(&professor)
}

func (p *ProfessorRepositoryImple) FindAll() []models.Professor {
	var professores []models.Professor
	p.Db.Find(&professores)
	return professores
}

func (p *ProfessorRepositoryImple) FindById(professorId uint) (models.Professor, error) {
	var professor models.Professor
	err := p.Db.First(&professor, professorId).Error
	return professor, err
}

func (p *ProfessorRepositoryImple) Save(professor models.Professor) {
	p.Db.Create(&professor)
}

func (p *ProfessorRepositoryImple) Update(professor models.Professor) {
	var updateProfessor = request.AtualizaProfessorRequest{
		Id:    professor.Id,
		Nome:  professor.Nome,
		Email: professor.Email,
	}
	p.Db.Model(&professor).Updates(updateProfessor)

}
