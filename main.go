package main

import (
	"controle-notas/src/configuration/database"
	"controle-notas/src/controller"
	"controle-notas/src/models"
	"controle-notas/src/repository"
	"controle-notas/src/router"
	"controle-notas/src/service/professor"
	"log"
	"net/http"

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
	//db.Table("professor").AutoMigrate(&models.Professor{})
	db.AutoMigrate(&models.Professor{})

	professorRepository := repository.NewProfessorRepositoryImple(db)

	professorService := professor.NewProfessorServiceImple(professorRepository, validate)

	professorController := controller.NewProfessorController(professorService)

	routes := router.NewRouter(professorController)

	server := &http.Server{
		Addr:    ":3000",
		Handler: routes,
	}

	server.ListenAndServe()

}
