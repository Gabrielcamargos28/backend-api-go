package data

type AtualizarProfessorRequest struct {
	Id    uint   `json:"id" validate:required`
	Nome  string `json:"nome" validate:"required"`
	Email string `json:"email" validate:"required"`
	CPF   string `json:"cpf"  validate:"required"`
}
