package controller

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/data"
	"controle-notas/src/service/atividade"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AtividadeController struct {
	AtividadeService atividade.AtividadeService
}

func NewAtividadeController(service atividade.AtividadeService) *AtividadeController {
	return &AtividadeController{
		AtividadeService: service,
	}
}

func (controller *AtividadeController) Create(ctx *gin.Context) {
	var criarRequisicao data.AtividadeRequest
	if err := ctx.ShouldBindJSON(&criarRequisicao); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.AtividadeService.Create(criarRequisicao)
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

func (controller *AtividadeController) Update(ctx *gin.Context) {
	atividadeId := ctx.Param("atividadeId")
	id, err := strconv.ParseUint(atividadeId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var requisicaoAtualizar data.AtualizarAtividadeRequest
	if err := ctx.ShouldBindJSON(&requisicaoAtualizar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requisicaoAtualizar.Id = uint(id)

	err = controller.AtividadeService.Update(requisicaoAtualizar)
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

func (controller *AtividadeController) Delete(ctx *gin.Context) {
	atividadeId := ctx.Param("atividadeId")
	id, err := strconv.ParseUint(atividadeId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = controller.AtividadeService.Delete(uint(id))
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

func (controller *AtividadeController) FindById(ctx *gin.Context) {
	atividadeId := ctx.Param("atividadeId")
	id, err := strconv.ParseUint(atividadeId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	atividadeResponse, err := controller.AtividadeService.FindById(uint(id))
	if err != nil {
		controller.handleRestErr(ctx, err)
		return
	}

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   atividadeResponse,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

/*
	func (controller *AtividadeController) FindAll(ctx *gin.Context) {
		atividadeResponse, err := controller.AtividadeService.FindAll()
		if err != nil {
			controller.handleRestErr(ctx, err)
			return
		}

		webResponse := data.ResponseApi{
			Code:   http.StatusOK,
			Status: "Ok",
			Data:   atividadeResponse,
		}
		ctx.JSON(http.StatusOK, webResponse)
	}
*/
func (controller *AtividadeController) FindAll(ctx *gin.Context) {
	atividadeResponse, err := controller.AtividadeService.FindAll()
	if err != nil {
		controller.handleRestErr(ctx, err)
		return
	}

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   atividadeResponse,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AtividadeController) handleRestErr(ctx *gin.Context, err error) {
	statusCode := http.StatusInternalServerError
	if restErr, ok := err.(*rest_err.RestErr); ok {
		statusCode = restErr.Campo
	}
	ctx.JSON(statusCode, gin.H{"error": err.Error()})
}
