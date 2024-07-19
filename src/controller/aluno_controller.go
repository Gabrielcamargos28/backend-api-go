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

func NewAlunoController(service aluno.AlunoService) *AlunoController {
	return &AlunoController{
		AlunoService: service,
	}
}

func (controller *AlunoController) Create(ctx *gin.Context) {
	var criarRequisicao data.AlunoRequest
	if err := ctx.ShouldBindJSON(&criarRequisicao); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := controller.AlunoService.Create(criarRequisicao); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Mensagem})
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

func (controller *AlunoController) Update(ctx *gin.Context) {
	alunoId := ctx.Param("alunoId")
	id, err := strconv.ParseUint(alunoId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var requisicaoAtualizar data.AtualizarAlunoRequest
	if err := ctx.ShouldBindJSON(&requisicaoAtualizar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requisicaoAtualizar.Id = uint(id)

	if err := controller.AlunoService.Update(requisicaoAtualizar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Mensagem})
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

func (controller *AlunoController) Delete(ctx *gin.Context) {
	alunoId := ctx.Param("alunoId")
	id, err := strconv.ParseUint(alunoId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := controller.AlunoService.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Mensagem})
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

func (controller *AlunoController) FindById(ctx *gin.Context) {
	alunoId := ctx.Param("alunoId")
	id, err := strconv.ParseUint(alunoId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	alunoResponse, restErr := controller.AlunoService.FindById(uint(id))
	if restErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": restErr.Mensagem})
		return
	}

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   alunoResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AlunoController) FindAll(ctx *gin.Context) {
	alunoResponse, restErr := controller.AlunoService.FindAll()
	if restErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": restErr.Mensagem})
		return
	}

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   alunoResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
