package ad

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Ok  bool `json:"ok"`
	Error string `json:"error,omitempty"`
}

func NewGroupHandler(context echo.Context) error {
	err := CreateNewGroup(context.Param("group_name"))
	if err != nil {
		response := &Response{ Ok: false, Error: err.Error()}
		return context.JSON(http.StatusOK, response)
	}
	return context.JSON(http.StatusOK, &Response{ Ok: true })
}