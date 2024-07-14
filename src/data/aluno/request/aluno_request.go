package request

type AlunoRequest struct {
	Nome      string `json:"nome" validate:"required"`
	Matricula string `json:"matricula" validate:"required"`
}
