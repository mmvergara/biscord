package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/mmvergara/biscord/go-backend/config"
	model "github.com/mmvergara/biscord/go-backend/models"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// Middleware decodes the share session cookie and packs the session into context
func (h *Handler) AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("jwt")

			// Allow unauthenticated users in
			if c == nil {
				next.ServeHTTP(w, r)
				return
			}

			if err != nil {
				SendErrorResponse(w, http.StatusForbidden, "Invalid cookie")
				return
			}

			token, err := validateJWT(c.Value)
			if err != nil {
				log.Println(err)
				SendErrorResponse(w, http.StatusForbidden, "Invalid token")
				return
			}

			userID, err := h.getUserIDFromToken(token)
			if err != nil {
				log.Println(err)
				SendErrorResponse(w, http.StatusForbidden, "Invalid token")
				return
			}

			// get the user from the database
			user, err := h.userRepo.GetUserByID(userID)
			if err != nil {
				log.Println(err)
				SendErrorResponse(w, http.StatusForbidden, err.Error())
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

func (a *Handler) getUserIDFromToken(token *jwt.Token) (model.UserID, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	nilUUID := uuid.UUID{}

	if !ok {
		return nilUUID, fmt.Errorf("invalid token claims")
	}

	userid, ok := claims["sub"].(string)
	if !ok {
		return nilUUID, fmt.Errorf("invalid token claims")
	}
	uid, err := uuid.Parse(userid)
	if err != nil {
		return nilUUID, fmt.Errorf("invalid token claims")
	}

	return uid, nil
}

func (h *Handler) createToken(userid model.UserID, username string, displayName string) (string, time.Time, error) {
	expires := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":         userid.String(),
			"username":    username,
			"displayName": displayName,
			"exp":         expires.Unix(),
		})

	tokenString, err := token.SignedString([]byte(config.Envs.JWTSecret))
	if err != nil {
		log.Println(err)
		log.Println(err)
		log.Println(err)
		return "", expires, err
	}

	return tokenString, expires, nil
}
