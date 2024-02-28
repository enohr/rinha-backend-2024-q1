package clientes

import "time"

type Transacao struct {
	Valor       int
	Tipo        string
	Descricao   string
	RealizadaEm time.Time
}
