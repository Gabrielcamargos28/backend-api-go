package data

type RemoverAlunoTurmaRequest struct {
	TurmaId  uint   `json:"turma_id" validate:"required"`
	AlunosId []uint `json:"alunos_id" validate:"required,dive,required"`
}
