package serializers

type ClientesTransacoesInput struct {
	Valor     int    `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

type ClientesTransacoesOutput struct {
	Limite int `json:"limite"`
	Saldo  int `json:"saldo"`
}
