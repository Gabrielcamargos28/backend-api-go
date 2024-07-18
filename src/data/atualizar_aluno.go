package data

type AtualizarAlunoRequest struct {
	Id        uint   `json:"id" validate:"required"`
	Nome      string `json:"nome" validate:"required"`
	Matricula string `json:"matricula" validate:"required"`
}
