package repository

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data"
	"controle-notas/src/models"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type NotaRepositoryImple struct {
	Db *gorm.DB
}

func NewNotaRepositoryImple(Db *gorm.DB) NotaRepository {
	return &NotaRepositoryImple{Db: Db}
}

func (n *NotaRepositoryImple) Delete(notaId uint) *rest_err.RestErr {
	var nota models.Nota

	result := n.Db.Where("id = ?", notaId).First(&nota)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return rest_err.NewNotFoundError("Nota não encontrada")
		}
		return rest_err.NewInternalServerError("Erro ao buscar nota")
	}

	result = n.Db.Delete(&nota)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "violates foreign key constraint") {
			return rest_err.NewBadRequestError("Não é possível deletar a nota pois ela está associada a uma ou mais turmas")
		}
		return rest_err.NewInternalServerError("Erro ao deletar nota")
	}

	return nil
}
func (n *NotaRepositoryImple) FindAll() ([]models.Nota, *rest_err.RestErr) {
	var notas []models.Nota
	resultado := n.Db.Preload("Aluno").Preload("Atividade.Turma").Find(&notas)
	if resultado.Error != nil {
		return nil, rest_err.NewInternalServerError("Erro ao buscar notas")
	}
	fmt.Println(notas, resultado)
	return notas, nil
}
func (r *NotaRepositoryImple) FindById(notaId uint) (models.Nota, *rest_err.RestErr) {
	var nota models.Nota
	if err := r.Db.Preload("Aluno").Preload("Atividade.Turma").First(&nota, notaId).Error; err != nil {
		if err != nil {
			return nota, rest_err.NewNotFoundError("Nota não encontrada")
		}
		return nota, rest_err.NewInternalServerError("Erro ao buscar nota")
	}
	return nota, nil
}

func (n *NotaRepositoryImple) Save(nota models.Nota) *rest_err.RestErr {
	if result := n.Db.Create(&nota); result.Error != nil {
		return rest_err.NewInternalServerError("Erro ao salvar atividade")
	}
	return nil
}

func (n *NotaRepositoryImple) Update(nota models.Nota) *rest_err.RestErr {
	var notaUpdate = data.AtualizarNota{
		Id:    nota.Id,
		Valor: nota.Valor,
	}
	if result := n.Db.Model(&nota).Updates(notaUpdate); result.Error != nil {
		return rest_err.NewInternalServerError("Erro ao atualizar nota")
	}
	return nil
}
