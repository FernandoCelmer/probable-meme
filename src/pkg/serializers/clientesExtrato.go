package serializers

import "time"

type UltimasTransacoes struct {
	Valor       int       `json:"valor"`
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	RealizadaEm time.Time `json:"realizada_em"`
}

type ClientesExtratoOutput struct {
	Saldo struct {
		Total       int       `json:"total"`
		DataExtrato time.Time `json:"data_extrato"`
		Limite      int       `json:"limite"`
	} `json:"saldo"`
	UltimasTransacoes []UltimasTransacoes `json:"ultimas_transacoes"`
}
