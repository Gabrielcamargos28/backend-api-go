package controller

import (
	"controle-notas/src/data"
	"controle-notas/src/data/turma/request"
	"controle-notas/src/service/turma"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TurmaController struct {
	TurmaService turma.TurmaService
}

func NewTurmaController(service turma.TurmaService) *TurmaController {
	return &TurmaController{
		TurmaService: service,
	}
}

func (controller *TurmaController) Create(ctx *gin.Context) {

	var criarRequisicao request.TurmaRequest
	if err := ctx.ShouldBindJSON(&criarRequisicao); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.TurmaService.Create(criarRequisicao)

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   criarRequisicao,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TurmaController) Update(ctx *gin.Context) {

	turmaId := ctx.Param("turmaId")
	id, err := strconv.ParseUint(turmaId, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var requisicaoAtualizar = request.AtualizaTurmaRequest{}
	if err := ctx.ShouldBindJSON(&requisicaoAtualizar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requisicaoAtualizar.Id = uint(id)

	controller.TurmaService.Update(requisicaoAtualizar)

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   requisicaoAtualizar,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TurmaController) Delete(ctx *gin.Context) {

	turmaId := ctx.Param("turmaId")

	id, err := strconv.ParseUint(turmaId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	controller.TurmaService.Delete(uint(id))

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TurmaController) FindById(ctx *gin.Context) {

	turmaId := ctx.Param("turmaId")
	id, err := strconv.ParseUint(turmaId, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	turmaResponse := controller.TurmaService.FindById(uint(id))

	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   turmaResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TurmaController) FindAll(ctx *gin.Context) {

	turmaResponse := controller.TurmaService.FindAll()
	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   turmaResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
func (controller *TurmaController) AdicionarAlunos(ctx *gin.Context) {
	var requisicao request.AdicioanrAlunosTurma
	if err := ctx.ShouldBindJSON(&requisicao); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro: ": err.Error()})
		return
	}
	controller.TurmaService.AdicionarAlunos(requisicao)
	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   requisicao,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
func (tc *TurmaController) RemoveAlunoDaTurma(ctx *gin.Context) {
	var requisicao request.RemoverAlunoTurmaRequest
	if err := ctx.ShouldBindJSON(&requisicao); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, alunoId := range requisicao.AlunosId {
		err := tc.TurmaService.RemoveAlunoTurma(alunoId, requisicao.TurmaId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao remover aluno da turma"})
			return
		}
	}
	webResponse := data.ResponseApi{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   requisicao,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
