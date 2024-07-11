package professor

type Professor struct {
	Nome  string `json:"nome"  binding: "required,min=4,max=50"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
}
