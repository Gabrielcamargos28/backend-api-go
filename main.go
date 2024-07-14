package main

import (
	"controle-notas/src/configuration/database"
	"controle-notas/src/controller"
	"controle-notas/src/models"
	"controle-notas/src/repository"
	"controle-notas/src/router"
	"controle-notas/src/service/aluno"
	"controle-notas/src/service/professor"
	"controle-notas/src/service/turma"
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

	db.AutoMigrate(&models.Professor{}, &models.Turma{}, &models.Aluno{})

	validate := validator.New()

	professorRepository := repository.NewProfessorRepositoryImple(db)
	turmaRepository := repository.NewTurmaRepositoryImple(db)
	alunoRepository := repository.NewAlunoRepositoryImple(db)

	professorService := professor.NewProfessorServiceImple(professorRepository, validate)
	turmaService := turma.NewTurmaServiceImple(turmaRepository, alunoRepository, validate)
	alunoService := aluno.NewAlunoServiceImple(alunoRepository, validate)

	professorController := controller.NewProfessorController(professorService)
	turmaController := controller.NewTurmaController(turmaService)
	alunoController := controller.NewAlunoController(alunoService)

	routes := router.NewRouter(professorController, turmaController, alunoController)

	server := &http.Server{
		Addr:    ":3000",
		Handler: routes,
	}

	log.Println("Servidor iniciado na porta 3000...")
	log.Fatal(server.ListenAndServe())
}