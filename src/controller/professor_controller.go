package controller

import (
	"controle-notas/src/data"
	"controle-notas/src/data/professor/request"
	"controle-notas/src/service/professor"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProfessorController struct {
	ProfessorService professor.ProfessorService
}

func NewProfessorController(sevice professor.ProfessorService) *ProfessorController {
	return &ProfessorController{
		ProfessorService: sevice,
	}
}

func (controller *ProfessorController) Create(ctx *gin.Context) {

	criarRequisicao := request.ProfessorRequest{}

	controller.ProfessorService.Create(criarRequisicao)

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProfessorController) Update(ctx *gin.Context) {

	requisicaoAtualizar := request.AtualizaProfessorRequest{}

	professorId := ctx.Param("professorId")
	id, err := strconv.Atoi(professorId)
	err.Error()
	requisicaoAtualizar.Id = id

	controller.ProfessorService.Update(requisicaoAtualizar)

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProfessorController) Delete(ctx *gin.Context) {

	professorId := ctx.Param("professorId")
	id, err := strconv.Atoi(professorId)
	err.Error()
	controller.ProfessorService.Delete(id)

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProfessorController) FindById(ctx *gin.Context) {

	professorId := ctx.Param("professorId")
	id, err := strconv.Atoi(professorId)
	err.Error()

	professorResponse := controller.ProfessorService.FindById(id)

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   professorResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProfessorController) FindAll(ctx *gin.Context) {

	professorResponse := controller.ProfessorService.FindAll()
	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   professorResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
