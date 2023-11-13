package models

type Loja struct {
	Id                int    `json:"idLoja"`
	IdEstabelecimento int    `json:"idEstabelecimento"`
	Nome              string `json:"nomeLoja"`
	RazaoSocial       string `json:"razaoSocialLoja"`
	Endereco          string `json:"enderecoLoja"`
	Estado            string `json:"estadoLoja"`
	Cidade            string `json:"cidadeLoja"`
	Cep               string `json:"cepLoja"`
}
