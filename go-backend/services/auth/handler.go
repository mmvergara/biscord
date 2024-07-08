package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/mmvergara/biscord/go-backend/services/repo"
)

type Handler struct {
	userRepo *repo.UserRepo
}

func NewHandler(userRepo *repo.UserRepo) *Handler {
	return &Handler{
		userRepo: userRepo,
	}
}

func (h *Handler) SignIn(c echo.Context) error {

	return c.HTML(200, "Sign In")
}
