package controller

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data"
	"controle-notas/src/service/nota"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotaController struct {
	NotaService nota.NotaService
}

func NewNotaController(service nota.NotaService) *NotaController {
	return &NotaController{
		NotaService: service,
	}
}

func (controller *NotaController) Create(ctx *gin.Context) {
	var criarRequisicao data.NotaRequest
	if err := ctx.ShouldBindJSON(&criarRequisicao); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.NotaService.Create(criarRequisicao)
	if err != nil {
		controller.handleRestErr(ctx, err)
		return
	}

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   criarRequisicao,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *NotaController) Update(ctx *gin.Context) {
	notaId := ctx.Param("notaId")
	id, err := strconv.ParseUint(notaId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var requisicaoAtualizar data.AtualizarNota
	if err := ctx.ShouldBindJSON(&requisicaoAtualizar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requisicaoAtualizar.Id = uint(id)

	err = controller.NotaService.Update(requisicaoAtualizar)
	if err != nil {
		controller.handleRestErr(ctx, err)
		return
	}

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   requisicaoAtualizar,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *NotaController) Delete(ctx *gin.Context) {
	notaId := ctx.Param("notaId")
	id, err := strconv.ParseUint(notaId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = controller.NotaService.Delete(uint(id))
	if err != nil {
		controller.handleRestErr(ctx, err)
		return
	}

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *NotaController) FindById(ctx *gin.Context) {
	notaId := ctx.Param("notaId")
	id, err := strconv.ParseUint(notaId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	notaResponse, err := controller.NotaService.FindById(uint(id))
	if err != nil {
		controller.handleRestErr(ctx, err)
		return
	}

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   notaResponse,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *NotaController) FindAll(ctx *gin.Context) {
	notaResponse, err := controller.NotaService.FindAll()
	if err != nil {
		controller.handleRestErr(ctx, err)
		return
	}

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   notaResponse,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *NotaController) handleRestErr(ctx *gin.Context, err error) {
	statusCode := http.StatusInternalServerError
	if restErr, ok := err.(*rest_err.RestErr); ok {
		statusCode = restErr.Campo
	}
	ctx.JSON(statusCode, gin.H{"error": err.Error()})
}
