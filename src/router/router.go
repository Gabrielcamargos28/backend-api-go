package router

import (
	"controle-notas/src/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(professorController *controller.ProfessorController) *gin.Engine {
	ctx.JSON(http.StatusOK, "")
}
