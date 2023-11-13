package rotas

import (
	"snet-apiGo/src/controllers/handlers"

	"github.com/labstack/echo"
)

func Rotas(e *echo.Echo) {

	//estabelecimento
	e.POST("/estabelecimento", handlers.ReceberRequisicaoCriarEstabelecimento)
	e.GET("/estabelecimentos", handlers.RecebeberRequisicaoListarEstabelecimentos)
	e.GET("/estabelecimentos/:id", handlers.RecebeberRequisicaoGetEstabelecimentoPorID)
	e.DELETE("/estabelecimentos/:id", handlers.RecebeberRequisicaoDeletarEstabelecimento)
	e.PUT("/estabelecimentos/:id", handlers.ReceberRequisicaoAtualizarEstabelecimento)
	e.GET("/estabelecimentos/lojas/:id", handlers.ReceberRequisicaoListarLojasDoEstabelecimento)

	//loja
	
	e.POST("/lojas", handlers.ReceberRequisicaoCriarLoja)
	e.GET("/lojas", handlers.RecebeberRequisicaoListarLojas)
	e.DELETE("/lojas/:id", handlers.RecebeberRequisicaoDeletarLoja)
	e.PUT("/lojas/:loja_id", handlers.ReceberRequisicaoAtualizarLoja)

}
