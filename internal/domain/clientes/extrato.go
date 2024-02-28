package clientes

import "time"

type Extrato struct {
	Saldo      Saldo
	Transacoes []Transacao
}

type Saldo struct {
	Total        int
	Data_extrato time.Time
	Limite       int
}
