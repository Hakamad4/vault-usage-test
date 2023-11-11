package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		ReadVaultConfigs()
		return c.String(http.StatusOK, "Olá, mundo com Echo!")
	})

	// Inicia o servidor na porta 8080
	e.Logger.Fatal(e.Start(":8080"))

}
