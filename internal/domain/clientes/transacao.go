package clientes

import "time"

type Transacao struct {
	Valor       int       `json:"valor"`
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	RealizadaEm time.Time `json:"realizada_em"`
}

// TODO: Rename to a better name
type TransacaoResponse struct {
	Limite int `json:"limite"`
	Saldo  int `json:"saldo"`
}
