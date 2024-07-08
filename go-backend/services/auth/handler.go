package auth

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	model "github.com/mmvergara/biscord/go-backend/models"
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

func (h *Handler) sendErrorResponse(c echo.Context, status int, message string) error {
	log.Println(message)
	return c.JSON(status, struct {
		Error string `json:"error"`
	}{
		Error: message,
	})
}

func (h *Handler) SignIn(c echo.Context) error {

	// Get user by email

	// Compare password

	// Generate JWT

	// Set JWT in cookie as httponly

	// Return sub,exp,username,displayName,avatarURL

	return c.HTML(200, "Sign In")
}

func (h *Handler) SignUp(c echo.Context) error {
	userForm := new(model.SignUpForm)
	err1 := c.Bind(&userForm)
	if err1 != nil {
		log.Println(err1)
		return h.sendErrorResponse(c, http.StatusBadRequest, "Failed to bind request body")
	}

	// Check if username is already taken
	yes, err := h.userRepo.DoesUsernameExist(userForm.Username)
	if err != nil {
		log.Println(err)
		return h.sendErrorResponse(c, http.StatusInternalServerError, "Failed to check if username exists")
	}

	if yes {
		return h.sendErrorResponse(c, http.StatusBadRequest, "Username is already taken")
	}

	// Check if email is already taken
	yes, err = h.userRepo.DoesEmailExist(userForm.Email)
	if err != nil {
		log.Println(err)
		return h.sendErrorResponse(c, http.StatusInternalServerError, "Failed to check if email exists")
	}
	if yes {
		return h.sendErrorResponse(c, http.StatusBadRequest, "Email is already taken")
	}

	// Hash password
	hashedPassword, err2 := HashPassword(userForm.Password)
	if err2 != nil {
		log.Println(err2)
		return h.sendErrorResponse(c, http.StatusInternalServerError, "Failed to hash password")
	}

	err3 := h.userRepo.CreateUser(*userForm, hashedPassword)
	if err3 != nil {
		log.Println(err3)
		return h.sendErrorResponse(c, http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusOK, &struct{}{})
}

func (h *Handler) SignOut(c echo.Context) error {
	return c.HTML(200, "Sign Out")
}
