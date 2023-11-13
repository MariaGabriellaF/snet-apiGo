package models

type Estabelecimento struct {
	Id                    int    `json:"id"`
	Nome                  string `json:"nome"`
	RazaoSocial           string `json:"razaoSocial"`
	Endereco              string `json:"endereco"`
	Estado                string `json:"estado"`
	Cidade                string `json:"cidade"`
	Cep                   string `json:"cep"`
	NumeroEstabelecimento int    `json:"numeroEst"`
	Lojas []Loja
}