package rotas

import "github.com/gin-gonic/gin"

func IniciarRotas(r *gin.RouterGroup) {
	r.GET("/usuario/retornaIdUsuario/:usuarioId")
	r.GET("/usuario/retornaIdUsuario/:usuarioEmail")
	r.POST("/usuario/criar")
	r.PUT("/usuario/editar/:usuarioId")
	r.DELETE("/usuario/deletar/:usuarioId")
}
