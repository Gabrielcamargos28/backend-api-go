package rest_err

import "net/http"

type RestErr struct {
	Mensagem string   `json: "mensagem"`
	Err      string   `json: "error`
	Codigo   int64    `json: "codigo`
	Causas   []Causas `json: "causa"`
}

type Causas struct {
	Campo    string `json: "campo`
	Mensagem string `json: "mensagem"`
}

func NewRestErr(mensagem string, err string, codigo int64, causas []Causas) *RestErr {
	return &RestErr{
		Mensagem: mensagem,
		Err:      err,
		Codigo:   codigo,
		Causas:   causas,
	}
}

func NewBadRequestError(mensagem string) *RestErr {
	return &RestErr{
		Mensagem: mensagem,
		Err:      "Pedido invalido",
		Codigo:   http.StatusBadRequest,
	}
}
func NewBadValidationError(mensagem string, causas []Causas) *RestErr {
	return &RestErr{
		Mensagem: mensagem,
		Err:      "Pedido invalido",
		Codigo:   http.StatusBadRequest,
		Causas:   causas,
	}
}
func NewInternalServerError(mensagem string, causas []Causas) *RestErr {
	return &RestErr{
		Mensagem: mensagem,
		Err:      "Erro interno",
		Codigo:   http.StatusInternalServerError,
		Causas:   causas,
	}
}
func NewUnauthorizedRequestError(mensagem string) *RestErr {
	return &RestErr{
		Mensagem: mensagem,
		Err:      "não autorizado",
		Codigo:   http.StatusUnauthorized,
	}
}
func NewForbiddenError(mensagem string) *RestErr {
	return &RestErr{
		Mensagem: mensagem,
		Err:      "forbidden",
		Codigo:   http.StatusForbidden,
	}
}
