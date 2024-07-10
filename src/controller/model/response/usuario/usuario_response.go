package usuario

type UsuarioResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Nome  string `json:"nome"`
	Idade int8   `json:"idade"`
}
