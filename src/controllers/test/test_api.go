package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"snet-apiGo/src/controllers/handlers"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCriarEstabelecimento(t *testing.T) {

	// Criar uma instância do Echo para o teste
	e := echo.New()

	// Criar um corpo de requisição JSON para o teste
	requestBody := `{"nome": "Estabelecimento Teste", "endereco": "Rua Teste, 123"}`

	// Criar uma requisição HTTP simulada com o corpo
	req := httptest.NewRequest(http.MethodPost, "/criar-estabelecimento", bytes.NewBufferString(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// Criar um contexto do Echo para o teste
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Chamar a função que está sendo testada
	err := handlers.ReceberRequisicaoCriarEstabelecimento(c)

	// Verificar se não há erros
	assert.NoError(t, err)

	// Verificar o código de status HTTP
	assert.Equal(t, http.StatusCreated, rec.Code)

}
