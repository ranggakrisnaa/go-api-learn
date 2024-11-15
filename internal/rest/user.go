package rest

import (
	"context"
	"go-api-learn/domain"
	"go-api-learn/internal/rest/util"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	Store(context.Context, *domain.User) (*uuid.UUID, error)
}

type UserHandler struct {
	Service UserService
}

func NewUserHandler(e *echo.Echo, svc UserService) {
	handler := &UserHandler{
		Service: svc,
	}

	e.POST("/auth/login", handler.Store)
}

func (h *UserHandler) Store(c echo.Context) (err error) {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	var ok bool
	if ok, err = util.ValidateStruct(&user); !ok {
		util.ErrorResponse(c, http.StatusBadRequest, "Invalid validation", err)
	}

	// ctx := c.Request().Context()

	return nil
}
