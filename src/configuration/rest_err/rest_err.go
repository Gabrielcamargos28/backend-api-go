package rest_err

import "net/http"

type RestErr struct {
	Mensagem string   `json: "mensagem"`
	Err      string   `json: "error`
	Campo    int      `json: "campo`
	Causas   []Causas `json: "causa"`
}

type Causas struct {
	Campo    string `json: "campo`
	Mensagem string `json: "mensagem"`
}

func (r *RestErr) Error() string {
	return r.Mensagem
}

func NewRestErr(mensagem string, err string, campo int, causas []Causas) *RestErr {
	return &RestErr{
		Mensagem: mensagem,
		Err:      err,
		Campo:    campo,
		Causas:   causas,
	}
}

func NewNotFoundError(mensagem string) *RestErr {
	return &RestErr{
		Mensagem: mensagem,
		Err:      "Não encontrado",
		Campo:    http.StatusNotFound,
	}
}

func NewBadRequestError(mensagem string) *RestErr {
	return &RestErr{
		Mensagem: mensagem,
		Err:      "Pedido invalido",
		Campo:    http.StatusBadRequest,
	}
}
func NewBadValidationError(mensagem string, causas []Causas) *RestErr {
	return &RestErr{
		Mensagem: mensagem,
		Err:      "Pedido invalido",
		Campo:    http.StatusBadRequest,
		Causas:   causas,
	}
}
func NewInternalServerError(mensagem string, causas []Causas) *RestErr {
	return &RestErr{
		Mensagem: mensagem,
		Err:      "Erro interno",
		Campo:    http.StatusInternalServerError,
		Causas:   causas,
	}
}
func NewUnauthorizedRequestError(mensagem string) *RestErr {
	return &RestErr{
		Mensagem: mensagem,
		Err:      "não autorizado",
		Campo:    http.StatusUnauthorized,
	}
}
func NewForbiddenError(mensagem string) *RestErr {
	return &RestErr{
		Mensagem: mensagem,
		Err:      "forbidden",
		Campo:    http.StatusForbidden,
	}
}
