package usuario

type UsuarioRequest struct {
	Email string `json:"email" binding: "required,email"`
	Senha string `json:"senha" binding: "required,min=3,max=8,containsany=!@#"`
	Nome  string `json:"nome"  binding: "required,min=4,max=50"`
	Idade int8   `json:"idade" binding: "required,min=1,max=140"`
}
