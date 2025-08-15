package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/neberson/pos-go-expert-fullcycle/tree/main/modulos/API/internal/dto"
	"github.com/neberson/pos-go-expert-fullcycle/tree/main/modulos/API/internal/entity"
	"github.com/neberson/pos-go-expert-fullcycle/tree/main/modulos/API/internal/infra/database"
)

type UserHandler struct {
	UserDB      database.UserInterface
	Jwt         *jwtauth.JWTAuth
	JwtExpireIn int
}

func NewUserHandler(userDB database.UserInterface, jwt *jwtauth.JWTAuth, jwtExpireIn int) *UserHandler {
	return &UserHandler{
		UserDB:      userDB,
		Jwt:         jwt,
		JwtExpireIn: jwtExpireIn,
	}
}

func (h *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJwtInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !u.IsValidPassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := h.Jwt.Encode(
		map[string]interface{}{
			"sub": u.ID.String(),
			"exp": time.Now().Add(time.Duration(h.JwtExpireIn) * time.Second).Unix(),
		})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
