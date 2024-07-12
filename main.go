package main

import (
	"controle-notas/src/configuration/database"
	"controle-notas/src/controller"
	"controle-notas/src/models"
	"controle-notas/src/repository"
	"controle-notas/src/service/professor"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao iniciar .env")
	}
	db := database.DatabaseConnection()
	validate := validator.New()
	db.Table("professor").AutoMigrate(&models.Professor{})

	//Repository
	professorRepository := repository.NewProfessorRepositoryImple(db)

	//Service
	professorService := professor.NewProfessorServiceImple(professorRepository, validate)

	//Controller
	professorController := controller.NewProfessorController(professorService)

	//Router
	routes := router.NewProfessorController(professorController)

	router := gin.Default()

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
