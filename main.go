package main

import (
	"snet-apiGo/src/controllers/rotas"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	rotas.Rotas(e)
	e.Start(":8080")
}
