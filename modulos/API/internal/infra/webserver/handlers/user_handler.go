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

type Error struct {
	Message string `json:"message"`
}

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

// GetJwt godoc
// @Summary Get JWT token for user authentication
// @Description Authenticate user and return JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.GetJwtInput true "user credentials"
// @Success 200 {object} dto.GetJwtOutput
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 500 {object} Error
// @Router /users/generation_token [post]
func (h *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJwtInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil || !u.IsValidPassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		error := Error{Message: "Invalid email or password"}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, tokenString, err := h.Jwt.Encode(
		map[string]interface{}{
			"sub": u.ID.String(),
			"exp": time.Now().Add(time.Duration(h.JwtExpireIn) * time.Second).Unix(),
		})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: "Failed to generate token"}
		json.NewEncoder(w).Encode(error)
		return
	}

	accessToken := dto.GetJwtOutput{AccessToken: tokenString}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
	w.WriteHeader(http.StatusOK)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserInput true "user request"
// @Success 201
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
