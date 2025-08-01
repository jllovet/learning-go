package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	ecomio "github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom/io"
	"github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom/services/auth"
	"github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom/types"
	"github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom/validation"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.RegisterUserPayload
	if err := ecomio.ParseJSON(r, &payload); err != nil {
		ecomio.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// validate the payload

	if err := validation.Validate().Struct(payload); err != nil {
		ecomio.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf("invlaid payload: %v", err.(validator.ValidationErrors)))
		return
	}

	// check if the user exists
	_, err := h.store.GetUserByEmail(payload.Email)

	// if err -> user already exists
	if err == nil {
		ecomio.WriteError(w, http.StatusBadRequest, fmt.Errorf("user already exists: %s", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		ecomio.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	// if the user does not exist, then create the new user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		ecomio.WriteError(w, http.StatusBadRequest, err)
		return
	}
	ecomio.WriteJSON(w, http.StatusCreated, nil)
}
