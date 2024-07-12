package request

type AtualizaProfessorRequest struct {
	Id    int    `json:"nome",validate:required`
	Nome  string `json:"nome",validate:"required"`
	Email string `json:"email",validate:"required"`
	CPF   string `json:"cpf",validate:"required"`
}
