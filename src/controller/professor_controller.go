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

func NewProfessorController(service professor.ProfessorService) *ProfessorController {
	return &ProfessorController{
		ProfessorService: service,
	}
}

func (controller *ProfessorController) Create(ctx *gin.Context) {
	var criarRequisicao request.ProfessorRequest
	if err := ctx.ShouldBindJSON(&criarRequisicao); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.ProfessorService.Create(criarRequisicao)
	if err != nil {
		ctx.JSON(err.Campo, gin.H{"error": err.Mensagem})
		return
	}

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   criarRequisicao,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProfessorController) Update(ctx *gin.Context) {
	professorId := ctx.Param("professorId")
	id, err := strconv.ParseUint(professorId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var requisicaoAtualizar = request.AtualizarProfessorRequest{}
	if err := ctx.ShouldBindJSON(&requisicaoAtualizar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requisicaoAtualizar.Id = uint(id)

	err = controller.ProfessorService.Update(requisicaoAtualizar)
	if err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "erro"})
		return
	}

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   requisicaoAtualizar,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProfessorController) Delete(ctx *gin.Context) {
	professorId := ctx.Param("professorId")
	id, err := strconv.ParseUint(professorId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = controller.ProfessorService.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
	id, err := strconv.ParseUint(professorId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	professorResponse, err := controller.ProfessorService.FindById(uint(id))
	if err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro"})
		return
	}

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   professorResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProfessorController) FindAll(ctx *gin.Context) {
	professorResponse, err := controller.ProfessorService.FindAll()
	if err != nil {
		ctx.JSON(err.Campo, gin.H{"error": err.Mensagem})
		return
	}

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   professorResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
