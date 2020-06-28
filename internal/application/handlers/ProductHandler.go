package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type ProductHandler struct {

}

func (h *ProductHandler) SaveProduct(c echo.Context) error {
	fmt.Print("HelloWorld")
	return nil
}

func (h *ProductHandler) FindProducts(c echo.Context) error {
	msg := "HelloWorld"
	return c.String(http.StatusOK, msg)
}


