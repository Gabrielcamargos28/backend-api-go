package main

import (
	"controle-notas/src/controller/rotas"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao iniciar .env")
	}

	router := gin.Default()
	grupoUsuario := router.Group("/usuario")
	grupoProfessor := router.Group("/professor")

	rotas.IniciarRotasUsuario(grupoUsuario)
	rotas.IniciarRotasProfessor(grupoProfessor)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
