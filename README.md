# snet-apiGo
A snet-apiGo oferece uma API RESTful que viabiliza o acesso para a criação e gerenciamento de estabelecimentos e lojas em um sistema.

**Recursos disponíveis na API:**
* Criação de Estabelecimento
* Visualização da Lista de Estabelecimentos Criados
* Visualização de um Único Estabelecimento por ID
* Atualização do Estabelecimento por ID
* Exclusão do Estabelecimento por ID
* Criação de Loja Vinculada a um Estabelecimento
* Visualização de Todas as Lojas Vinculadas a um Estabelecimento
* Visualização de uma Única Loja por ID
* Atualização da Loja por seu ID e o ID do Estabelecimento Associado
* Exclusão da Loja por seu ID e o ID do Estabelecimento Associado

# Linguagem e framework utilizado 

* **Linguagem:** Go (Golang)
* **Framework:** Echo

O Echo é um framework da linguagem Golang que facilita a criação de APIs, proporcionando uma estrutura robusta e simplificada para o desenvolvimento de aplicativos web.

Estas informações detalham a utilização do framework Echo juntamente com a linguagem Golang na construção da API NovoTeste.

## Métodos
Requisições para a API devem seguir os padrões:
| Método | Descrição |
|---|---|
| `GET` | Retorna informações de um ou mais registros. |
| `POST` | Utilizado para criar um novo registro. |
| `PUT` | Atualiza dados de um registro ou altera sua situação. |
| `DELETE` | Remove um registro do sistema. |

## Respostas

| Código | Descrição |
|---|---|
| `200` | Requisição executada com sucesso (success).|
| `201` | Requisição criada com sucesso (success).|
| `400` | Erros de validação ou os campos informados não existem no sistema.|
| `401` | Dados de acesso inválidos.|
| `404` | Registro pesquisado não encontrado (Not found).|
| `405` | Método não implementado.|
| `422` | Dados informados estão fora do escopo definido para o campo.|
| `500` | Internal Server Error.| ****

# Recursos da API

## Estabelecimento

## Criar estabelecimento (ReceberRequisicaoCriarEstabelecimento) [POST]

+ Atributos do estabelecimento
  
  + id: o id é gerado de forma automática por uma função chamada GerarID, que sorteia números randômicos. (int, obrigatório)
  + nome: nome do estabelecimento (string, obrigatório)
  + razaoSocial (string, obrigatório)
  + endereco (string, obrigatório)
  + estado (string, obrigatório)
  + cidade (string, obrigatório)
  + cep (string, obrigatório)
  + numeroEst: o numero do estabelecimento é gerado de forma automática por uma função chamada GerarID, que sorteia números randômicos. (int, obrigatório)
  + lojas (array, opicional)


    + Body

            {
              "nome": "",
              "razaoSocial": "",
              "endereco": "",
              "estado": "",
              "cidade": "",
              "cep": ""
            }
  
    + Response 201 (created/json)
      
            {
            	"id": 178,
            	"nome": "nome do estabelecimento",
            	"razaoSocial": "Razão Social",
            	"endereco": "Endereço do estabelecimento",
            	"estado": "Estado do estabelecimento",
            	"cidade": "Cidade do estabelecimento",
            	"cep": "CEP do estabelecimento",
            	"numeroEst": 44,
            	"lojas": null
            }

## Visualizar todos os estabelecimentos (RecebeberRequisicaoListarEstabelecimentos) [GET]

+ Request (application/json)

  + Response 200 (application/json)
    Mostra todos os estabelecimentos
    
    + Body
      
          	{
          		"id": 1,
          		"nome": "",
          		"razaoSocial": "",
          		"endereco": "",
          		"estado": "",
          		"cidade": "",
          		"cep": "",
          		"numeroEst": 2,
          		"lojas": []
          	}

 ## Visualizar um estabelecimento por ID (RecebeberRequisicaoGetEstabelecimentoPorID) [GET/estabelecimentos/id]

+ Request (application/json)

  + Response 200 (application/json)
      Mostrar o estabelecimento referente ao ID


      + Body
   
            {{
            	"nomeLoja": "gabi",
            	"razaoSocialLoja": "gabi",
            	"enderecoLoja": "gabi",
            	"estadoLoja": "g",
            	"cidadeLoja": "f",
            	"cepLoja": "f",
            	"numeroEstabelecimentoLoja": 200
            }
            	"id": 994,
            	"nome": "nome do estabelecimento",
            	"razaoSocial": "Razão Social",
            	"endereco": "Endereço do estabelecimento",
            	"estado": "Estado do estabelecimento",
            	"cidade": "Cidade do estabelecimento",
            	"cep": "CEP do estabelecimento",
            	"numeroEst": 780,
            	"lojas": null
            }

  + Response 404 (application/json)
      Quando registro não for encontrado.


      + Body

          {
          	"message": "Estabelecimento não encontrado"
            "message": "Not Found"
          }


## Deletar um estabelecimento por ID (RecebeberRequisicaoDeletarEstabelecimento) [DELETE/estabelecimentos/id]

+ Request (application/json)

  + Response 200 (application/json)
      Excluir o estabelecimento referente ao ID
    

      + Body
   
            "Estabelecimento excluído com sucesso"
        
      

  + Response 404 (application/json)
        Quando registro não for encontrado.



       + Body
   
            {
            	"message": "Não existe esse estabelecimento"
              "message": "Not Found"
            }

 
## Atualizar o estabelecimento por ID (ReceberRequisicaoAtualizarEstabelecimento) [PUT/estabelecimento/id]
Não é possível atualizar o ID e nem o número do estabelecimento

+ Request (application/json)

  + Response 200 (application/json)
      Atualziar o estabelecimento referente ao ID


      + Body
           
            {
              "nome": "NovoNome",
              "razaoSocial": "Nova Razão Social",
              "endereco": "Novo endereço",
              "estado": "Novo estado",
              "cidade": "Nova cidade",
              "cep": "Novo cep"
            }


  + Response 404 (application/json)
      Quando registro não for encontrado.


      + Body
          
              {
              	"message": "Estabelecimento não encontrado"
                "message": "Not Found"
              }

  + Response 422 (application/json)
      Erro de validação

      + Body
        
              {
              	"message": "Erro de validação dos dados"
              }


## Loja

## Criar Loja vinculada a um estabelecimento (ReceberRequisicaoCriarLojaELigarAUmEstabelecimento) [POST/idEstabelecimento/criarLoja]

+ Atributos da loja
  
  + id: o id é gerado de forma automática por uma função chamada GerarID, que sorteia números randômicos. (int, obrigatório)
  + nome: nome do estabelecimento (string, obrigatório)
  + razaoSocial (string, obrigatório)
  + endereco (string, obrigatório)
  + estado (string, obrigatório)
  + cidade (string, obrigatório)
  + cep (string, obrigatório)
  + numeroEst: o numero do estabelecimento é gerado de forma automática por uma função chamada GerarID, que sorteia números randômicos. (int, obrigatório)
  


    + Body

           {
            	"nomeLoja": "",
            	"razaoSocialLoja": "",
            	"enderecoLoja": "",
            	"estadoLoja": "",
            	"cidadeLoja": "",
            	"cepLoja": "",
            }

  
    + Response 201 (created/json)
      
            {
            	"id": 184,
            	"nome": "nome da loja",
            	"razaoSocial": "Razão Social",
            	"endereco": "Endereço da loja",
            	"estado": "Estado da loja",
            	"cidade": "Cidade da loja",
            	"cep": "CEP da loja",
            	"numeroEst": 12,
            	"lojas": null
            }



## Visualizar todos as lojas ligadas a um estabelecimento (ReceberRequisicaoListarLojasDoEstabelecimento) [GET/id/lojas]

+ Request (application/json)

  + Response 200 (application/json)
    Mostra todos as lojas ligadas a um estabelecimento
    
    + Body

                  [
        	{
        		"idLoja": 28,
        		"nomeLoja": "Loja do estabelecimento1",
        		"razaoSocialLoja": "Loja",
        		"enderecoLoja": "rua da loja",
        		"estadoLoja": "etado da loja",
        		"cidadeLoja": "cidade da loja",
        		"cepLoja": "000000",
        		"numeroEstabelecimentoLoja": 865
        	},
        	{
        		"idLoja": 343,
        		"nomeLoja": "Outra loja do estabelecimento1",
        		"razaoSocialLoja": "loja1",
        		"enderecoLoja": "rua da loja1",
        		"estadoLoja": "rua da loja1",
        		"cidadeLoja": "rua da loja1",
        		"cepLoja": "211111",
        		"numeroEstabelecimentoLoja": 865
        	}
        ]


## Deletar uma loja pelo seu ID (ReceberRequisicaoDeletarLoja) [DELETE/lojas/id]

+ Request (application/json)

  + Response 200 (application/json)
      Excluir a loja referente ao ID
    

      + Body
   
            "Loja excluída com sucesso"
        
      

  + Response 404 (application/json)
        Quando registro não for encontrado.


       + Body
   
            {
            	"message": "Não existe uma loja com esse ID"
               "message": "Not Found"
            }


## Atualizar a loja pelo seu ID e pelo ID do estabelecimento que ela está ligadad (ReceberRequisicaoAtualizarLoja) [PUT/estabelecimento/id/lojas/id]
Não é possível atualizar o ID

+ Request (application/json)

  + Response 200 (application/json)
      Atualziar a loja referente ao seu ID e ao estabelecimento que está ligado


      + Body
           
            {
        			"nomeLoja": "",
        			"razaoSocialLoja": "",
        			"enderecoLoja": "",
        			"estadoLoja": "",
        			"cidadeLoja": "",
        			"cepLoja": "",
        			"numeroEstabelecimentoLoja": 
        		}

  + Response 404 (application/json)
      Quando registro não for encontrado.


      + Body


              {
              	"message": "Loja não encontrada"
                "message": "Not Found"
              }
    

  + Response 422 (application/json)
      Erro de validação

      + Body
        
              {
              	"message": "Erro de validação dos dados"
              }
