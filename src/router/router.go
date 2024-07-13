package router

import (
	"controle-notas/src/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(professorController *controller.ProfessorController, turmaController *controller.TurmaController) *gin.Engine {

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
	turmaRouter.DELETE("/deletar/:turmaId", turmaController.Delete)

	return router
}
