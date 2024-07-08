package auth

import (
	"encoding/json"
	"log"
	"net/http"

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

type ErrorResponse struct {
	Error string `json:"error"`
}

type JsonResponse struct {
	Data interface{} `json:"data"`
}

func SendErrorResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	errorResponse := ErrorResponse{Error: message}
	json.NewEncoder(w).Encode(errorResponse)
}

func SendResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := JsonResponse{Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) SignInHandler(w http.ResponseWriter, r *http.Request) {
	signInForm := new(model.SignInForm)

	err := json.NewDecoder(r.Body).Decode(signInForm)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, http.StatusBadRequest, "Invalid Credentials")
		return
	}

	// Get user by email
	user, err := h.userRepo.GetUserByEmail(signInForm.Email)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// Compare password
	if !ComparePasswords(user.Password, signInForm.Password) {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid Credentials")
		return
	}

	// Generate JWT
	tokenString, exp, err := h.createToken(user.ID, user.Username, user.DisplayName)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Set JWT in cookie as httpOnly
	httpOnlyCookie := http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  exp,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &httpOnlyCookie)

	var res = &struct {
		UserID      model.UserID `json:"sub"`
		Username    string       `json:"username"`
		DisplayName string       `json:"displayName"`
		Exp         int64        `json:"exp"`
	}{
		UserID:      user.ID,
		Username:    user.Username,
		DisplayName: user.DisplayName,
	}

	SendResponse(w, http.StatusOK, res)
}

func (h *Handler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	userForm := new(model.SignUpForm)

	err := json.NewDecoder(r.Body).Decode(userForm)

	if err != nil {
		log.Println(err)
		SendErrorResponse(w, http.StatusBadRequest, "Failed to bind request body")
		return
	}

	// Check if username is already taken
	yes, err := h.userRepo.DoesUsernameExist(userForm.Username)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to check if username exists")
		return
	}

	if yes {
		SendErrorResponse(w, http.StatusBadRequest, "Username is already taken")
		return
	}

	// Check if email is already taken
	yes, err = h.userRepo.DoesEmailExist(userForm.Email)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to check if email exists")
		return
	}
	if yes {
		SendErrorResponse(w, http.StatusBadRequest, "Email is already taken")
		return
	}

	// Hash password
	hashedPassword, err := HashPassword(userForm.Password)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	err = h.userRepo.CreateUser(*userForm, hashedPassword)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	SendResponse(w, http.StatusCreated, nil)
	return
}

func (h *Handler) SignOutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})
	SendResponse(w, http.StatusOK, nil)
}
