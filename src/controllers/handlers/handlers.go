package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"snet-apiGo/src/models"
	"snet-apiGo/src/services"

	"strconv"

	"github.com/labstack/echo"
)

// Estabelecimento
func ReceberRequisicaoCriarEstabelecimento(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("(Handlers)Erro ao ler o corpo da solicitação:", err)
		return c.String(http.StatusBadRequest, "Erro ao ler o corpo da solicitação")
	}
	log.Println("(Handlers)Corpo da requisição: ", string(body))

	var estabelecimento models.Estabelecimento

	if err := json.Unmarshal(body, &estabelecimento); err != nil {
		log.Println("(Handlers)Erro ao decodificar o corpo da solicitação:", err)
		return c.String(http.StatusBadRequest, "Erro ao decodificar o corpo da solicitação")
	}

	log.Println("(Handlers)estabelecimento: ", estabelecimento)

	novoEstabelecimento, err := services.CriarEstabelecimento(c.Response().Writer, estabelecimento)
	if err != nil {
		log.Println("(Handlers)Erro ao criar o estabelecimento:", err)
		return c.String(http.StatusInternalServerError, "Erro ao criar o estabelecimento")
	}

	log.Println("(Handlers)Estabelecimento add número: ", estabelecimento)
	return c.JSON(http.StatusCreated, novoEstabelecimento)
}

func RecebeberRequisicaoListarEstabelecimentos(c echo.Context) error {
	estabelecimentos := services.GetEstabelecimentos(c.Response(), c.Request())
	log.Println("(Handlers)Estabelecimentos recuperados com sucesso")
	return c.JSON(http.StatusOK, estabelecimentos)
}

func RecebeberRequisicaoGetEstabelecimentoPorID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID inválido")
	}

	log.Printf("(Handlers)Recebendo requisição para buscar estabelecimento pelo ID: %d\n", id)

	estabelecimento, err := services.GetEstabelecimentoPorID(id)
	if err != nil {
		log.Printf("(Handlers)Erro ao buscar estabelecimento pelo ID %d: %s\n", id, err.Error())
		return echo.NewHTTPError(http.StatusNotFound, "Estabelecimento não encontrado")
	}
	log.Println("(Handlers)Estabelecimento recuperado:", estabelecimento)
	return c.JSON(http.StatusOK, estabelecimento)

}

func RecebeberRequisicaoDeletarEstabelecimento(c echo.Context) error {
	estabelecimentoID := c.Param("id")
	id, err := strconv.Atoi(estabelecimentoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID inválido")
	}

	err = services.DeletarEstabelecimentoPorID(id)
	if err != nil {
		log.Printf("(Handlers)Erro ao deletar estabelecimento pelo ID %s: %s\n", estabelecimentoID, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, "Estabelecimento excluído com sucesso")
}

func ReceberRequisicaoAtualizarEstabelecimento(c echo.Context) error {
	estabelecimentoID := c.Param("id")
	id, err := strconv.Atoi(estabelecimentoID)
	if err != nil {
		log.Println("(Handlers)Erro ao converter ID para inteiro:", err)
		return echo.NewHTTPError(http.StatusBadRequest, "ID inválido")
	}
	log.Printf("(Handlers)ID do estabelecimento: %d\n", id)

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("(Handlers)Erro ao ler a solicitação:", err)
		return c.String(http.StatusBadRequest, "Erro ao ler a solicitação")
	}

	var estabelecimento models.Estabelecimento
	if err := json.Unmarshal(body, &estabelecimento); err != nil {
		log.Println("(Handlers)Erro ao decodificar o corpo da solicitação:", err)
		return c.String(http.StatusBadRequest, "Erro ao decodificar o corpo da solicitação")
	}
	estabelecimento.Id = id

	err = services.AtualizarEstabelecimentoPorID(estabelecimento)

	return c.JSON(http.StatusOK, "Estabelecimento atualizado com sucesso")
}

// Loja

func ReceberRequisicaoCriarLoja(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("(Handlers)Erro ao ler o corpo da solicitação:", err)
		return c.String(http.StatusBadRequest, "Erro ao ler o corpo da solicitação")
	}
	log.Println("(Handlers)Corpo da requisição: ", string(body))

	var loja models.Loja
	if err := json.Unmarshal(body, &loja); err != nil {
		log.Println("(Handlers)Erro ao decodificar o corpo da solicitação:", err)
		return c.String(http.StatusBadRequest, "Erro ao decodificar o corpo da solicitação")
	}

	log.Println("(Handlers)loja: ", loja)

	novaLoja, err := services.CriarLoja(c.Response().Writer, loja)
	if err != nil {
		log.Println("(Handlers)Erro ao criar a loja:", err)
		return c.String(http.StatusInternalServerError, "Erro ao criar loja")
	}

	log.Println("(Handlers)Loja add número: ", loja)
	return c.JSON(http.StatusCreated, novaLoja)
	
}

func RecebeberRequisicaoListarLojas(c echo.Context) error {
	lojas := services.GetLojas(c.Response(), c.Request())
	log.Printf("(Handlers)Lojas recuperadas: %+v\n", lojas)
	return c.JSON(http.StatusOK, lojas)
}

func RecebeberRequisicaoDeletarLoja(c echo.Context) error {
    lojaID := c.Param("id")
    id, err := strconv.Atoi(lojaID)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "ID inválido")
    }

    err = services.DeletarLojaPorID(id)
    if err != nil {
        log.Printf("(Handlers)Erro ao deletar loja pelo ID %s: %s\n", lojaID, err.Error())
        return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao deletar loja")
    }

    return c.JSON(http.StatusOK, "Loja excluída com sucesso")
}

func ReceberRequisicaoAtualizarLoja(c echo.Context) error {
    lojaID := c.Param("id")
    id, err := strconv.Atoi(lojaID)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "ID inválido")
    }

    body, err := io.ReadAll(c.Request().Body)
    if err != nil {
        log.Println("(Handlers)Erro ao ler o corpo da solicitação:", err)
        return c.String(http.StatusBadRequest, "Erro ao ler o corpo da solicitação")
    }

    var loja models.Loja
    if err := json.Unmarshal(body, &loja); err != nil {
        log.Println("(Handlers)Erro ao decodificar o corpo da solicitação:", err)
        return c.String(http.StatusBadRequest, "Erro ao decodificar o corpo da solicitação")
    }
    loja.Id = id

    err = services.AtualizarLojaPorID(loja)
    if err != nil {
        log.Printf("(Handlers)Erro ao atualizar loja pelo ID %s: %s\n", lojaID, err.Error())
        return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao atualizar loja")
    }

    return c.JSON(http.StatusOK, "Loja atualizada com sucesso")
}

func ReceberRequisicaoListarLojasDoEstabelecimento(c echo.Context) error {
	estabelecimentoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID inválido")
	}
	log.Printf("(Handlers)ID do estabelecimento: %d\n", estabelecimentoID)

	estabelecimentos, err := services.ListarLojasDoEstabelecimentoPorID(estabelecimentoID)
	if err != nil {
		log.Println("Erro ao listar lojas do estabelecimento:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	log.Printf("(Handlers)Lojas do estabelecimento recuperadas")
	return c.JSON(http.StatusOK, estabelecimentos)
}