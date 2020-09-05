package handlers

import "github.com/labstack/echo/v4"

type ListHandler struct {

}

func CreateListHandler() ListHandler {
	return ListHandler{}
}

func (h *ListHandler) SaveList(echo.Context) error {

	return nil
}
