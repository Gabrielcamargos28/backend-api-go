package data

type ResponseApi struct {
	Code   int         `json:"codigo"`
	Status string      `json:"status"`
	Data   interface{} `json:"data.omitempty"`
}
