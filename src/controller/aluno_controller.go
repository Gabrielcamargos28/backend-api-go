package controller

import (
	"controle-notas/src/data"
	"controle-notas/src/service/aluno"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlunoController struct {
	AlunoService aluno.AlunoService
}

func NewAlunoController(sevice aluno.AlunoService) *AlunoController {
	return &AlunoController{
		AlunoService: sevice,
	}
}

func (controller *AlunoController) Create(ctx *gin.Context) {

	var criarRequisicao data.AlunoRequest
	if err := ctx.ShouldBindJSON(&criarRequisicao); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.AlunoService.Create(criarRequisicao)

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   criarRequisicao,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AlunoController) Update(ctx *gin.Context) {

	alunoId := ctx.Param("alunoId")
	id, err := strconv.ParseUint(alunoId, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var requisicaoAtualizar = data.AtualizarAlunoRequest{}
	if err := ctx.ShouldBindJSON(&requisicaoAtualizar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requisicaoAtualizar.Id = uint(id)

	controller.AlunoService.Update(requisicaoAtualizar)

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   requisicaoAtualizar,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AlunoController) Delete(ctx *gin.Context) {

	alunoId := ctx.Param("alunoId")

	id, err := strconv.ParseUint(alunoId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	controller.AlunoService.Delete(uint(id))

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AlunoController) FindById(ctx *gin.Context) {

	alunoId := ctx.Param("alunoId")
	id, err := strconv.ParseUint(alunoId, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	professorResponse := controller.AlunoService.FindById(uint(id))

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   professorResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AlunoController) FindAll(ctx *gin.Context) {

	professorResponse := controller.AlunoService.FindAll()
	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   professorResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
