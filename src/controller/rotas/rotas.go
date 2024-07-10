package rotas

import (
	"controle-notas/src/controller/professor"
	"controle-notas/src/controller/usuario"

	"github.com/gin-gonic/gin"
)

func IniciarRotasUsuario(r *gin.RouterGroup) {
	r.GET("/retornaIdUsuario/:usuarioId", usuario.BuscarUsuarioPorId)
	r.GET("/retornaEmailUsuario/:usuarioEmail", usuario.BuscarUsuarioPorEmail)
	r.POST("/criar", usuario.CriarUsuario)
	r.PUT("/editar/:usuarioId", usuario.AtualizarUsuario)
	r.DELETE("/deletar/:usuarioId", usuario.DeletarUsuario)
}

func IniciarRotasProfessor(r *gin.RouterGroup) {
	r.GET("/retornaIdProfessor/:professorId", professor.BuscarProfessorPorId)
	r.GET("/retornaEmailProfessor/:professorEmail", professor.BuscarProfessorPorEmail)
	r.POST("/criar", professor.CriarProfessor)
	r.PUT("/editar/:professorId", professor.AtualizarProfessor)
	r.DELETE("/deletar/:professorId", professor.DeletarProfessor)
}
