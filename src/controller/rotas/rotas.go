package rotas

import (
	"controle-notas/src/controller/usuario"

	"github.com/gin-gonic/gin"
)

func IniciarRotas(r *gin.RouterGroup) {
	r.GET("/usuario/retornaIdUsuario/:usuarioId", usuario.BuscarUsuarioPorId)
	r.GET("/usuario/retornaIdUsuario/:usuarioEmail", usuario.BuscarUsuarioPorEmail)
	r.POST("/usuario/criar", usuario.CriarUsuario)
	r.PUT("/usuario/editar/:usuarioId", usuario.AtualizarUsuario)
	r.DELETE("/usuario/deletar/:usuarioId", usuario.DeletarUsuario)
}
