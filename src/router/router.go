package router

import (
	"controle-notas/src/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(professorController *controller.ProfessorController) *gin.Engine {

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

	return router
}
