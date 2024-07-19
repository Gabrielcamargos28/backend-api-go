package router

import (
	"controle-notas/src/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(professorController *controller.ProfessorController, turmaController *controller.TurmaController, alunoController *controller.AlunoController, atividadeController *controller.AtividadeController, notaController *controller.NotaController) *gin.Engine {

	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Bem vindo")
	})
	baseRouter := router.Group("/api")

	professorRouter := baseRouter.Group("/professor")
	professorRouter.GET("/listarTodos", professorController.FindAll)
	professorRouter.GET("/listar/:professorId", professorController.FindById)
	professorRouter.POST("/criarProfessor", professorController.Create)
	professorRouter.PUT("/atualizar/:professorId", professorController.Update)
	professorRouter.DELETE("/deletar/:professorId", professorController.Delete)

	turmaRouter := baseRouter.Group("/turma")
	turmaRouter.GET("/listarTodos", turmaController.FindAll)
	turmaRouter.GET("/listar/:turmaId", turmaController.FindById)
	turmaRouter.POST("/criarTurma", turmaController.Create)
	turmaRouter.PUT("/atualizar/:turmaId", turmaController.Update)
	turmaRouter.PUT("/adicionarAlunos", turmaController.AdicionarAlunos)
	turmaRouter.PUT("/removerAluno", turmaController.RemoverAlunoTurma)
	turmaRouter.DELETE("/deletar/:turmaId", turmaController.Delete)
	turmaRouter.GET("/listar-atividades/:turmaId", turmaController.GetAtividadesByTurmaId)

	alunoRouter := baseRouter.Group("/aluno")
	alunoRouter.GET("/listarTodos", alunoController.FindAll)
	alunoRouter.GET("/listar/:alunoId", alunoController.FindById)
	alunoRouter.POST("/criarAluno", alunoController.Create)
	alunoRouter.PUT("/atualizar/:alunoId", alunoController.Update)
	alunoRouter.DELETE("/deletar/:alunoId", alunoController.Delete)

	atividadeRouter := baseRouter.Group("/atividade")
	atividadeRouter.GET("/listarTodos", atividadeController.FindAll)
	atividadeRouter.GET("/listar/:atividadeId", atividadeController.FindById)
	atividadeRouter.POST("/criarAtividade", atividadeController.Create)
	atividadeRouter.PUT("/atualizar/:atividadeId", atividadeController.Update)
	atividadeRouter.DELETE("/deletar/:atividadeId", atividadeController.Delete)

	notaRouter := baseRouter.Group("/nota")
	notaRouter.GET("/listarTodos", notaController.FindAll)
	notaRouter.GET("/listar/:notaId", notaController.FindById)
	notaRouter.POST("/criarNota", notaController.Create)
	notaRouter.PUT("/atualizar/:notaId", notaController.Update)
	notaRouter.DELETE("/deletar/:notaId", notaController.Delete)
	router.GET("/api/aluno/:alunoId/notas", notaController.FindNotasByAluno)

	return router
}
