package professor

type AtualizaProfessorRequest struct {
	Nome  string `json:"nome",validate:"required"`
	Email string `json:"email",validate:"required"`
	CPF   string `json:"cpf",validate:"required"`
}
