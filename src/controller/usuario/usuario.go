package usuario

import (
	"controle-notas/src/configuration/rest_err"
	"controle-notas/src/controller/model/request/usuario"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CriarUsuario(c *gin.Context) {
	var usuarioRequest usuario.UsuarioRequest

	if err := c.ShouldBindJSON(&usuarioRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Algo de errado aconteceu nos campos, error=%s", err.Error))
		c.JSON(restErr.Campo, restErr)

		return
	}
}
func BuscarUsuarioPorId(c *gin.Context) {

}
func BuscarUsuarioPorEmail(c *gin.Context) {

}
func DeletarUsuario(c *gin.Context) {

}
func AtualizarUsuario(c *gin.Context) {

}
