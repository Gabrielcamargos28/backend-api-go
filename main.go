package main

import (
	"controle-notas/src/configuration/database"
	"controle-notas/src/controller"
	"controle-notas/src/models"
	"controle-notas/src/repository"
	"controle-notas/src/router"
	"controle-notas/src/service/aluno"
	"controle-notas/src/service/atividade"
	"controle-notas/src/service/nota"
	"controle-notas/src/service/professor"
	"controle-notas/src/service/turma"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao iniciar .env")
	}

	db := database.DatabaseConnection()

	db.AutoMigrate(&models.Professor{}, &models.Turma{}, &models.Aluno{}, &models.Atividade{}, &models.Nota{}) // &models.AlunoNota{}

	validate := validator.New()

	professorRepository := repository.NewProfessorRepositoryImple(db)
	turmaRepository := repository.NewTurmaRepositoryImple(db)
	alunoRepository := repository.NewAlunoRepositoryImple(db)
	atividadeRepository := repository.NewAtividadeRepositoryImple(db)
	notaRepository := repository.NewNotaRepositoryImple(db)

	professorService := professor.NewProfessorServiceImple(professorRepository, validate)
	turmaService := turma.NewTurmaServiceImple(turmaRepository, alunoRepository, validate)
	alunoService := aluno.NewAlunoServiceImple(alunoRepository, validate)
	atividadeService := atividade.NewAtividadeServiceImple(atividadeRepository, turmaRepository, validate)
	notaService := nota.NewNotaServiceImple(notaRepository, atividadeRepository, validate)

	professorController := controller.NewProfessorController(professorService)
	turmaController := controller.NewTurmaController(turmaService)
	alunoController := controller.NewAlunoController(alunoService)
	atividadeController := controller.NewAtividadeController(atividadeService)
	notaController := controller.NewNotaController(notaService)

	routes := router.NewRouter(professorController, turmaController, alunoController, atividadeController, notaController)

	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},                             // Origens permitidas
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},  // Métodos permitidos
		AllowedHeaders:   []string{"Content-Type", "Authorization"}, // Cabeçalhos permitidos
		AllowCredentials: true,
	})

	corsHandler := corsOptions.Handler(routes)

	server := &http.Server{
		Addr:    ":3000",
		Handler: corsHandler,
	}

	log.Println("Servidor iniciado na porta 3000...")
	log.Fatal(server.ListenAndServe())
}
