package professor

type Professor struct {
	ID    uint   `json: "id"`
	Nome  string `json:"nome"`
	Email string `json:"email" binding: "required,email"`
	CPF   string `json:"cpf"`
}
