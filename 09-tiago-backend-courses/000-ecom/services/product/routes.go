package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	ecomio "github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom/io"
	"github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom/types"
	"github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom/validation"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/product", h.handleGetProducts).Methods("GET")
	router.HandleFunc("/product/{id:[0-9]+}", h.handleGetProductById).Methods("GET")
	router.HandleFunc("/product", h.handleCreateProduct).Methods("POST")
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()
	if err != nil {
		ecomio.WriteError(w, http.StatusInternalServerError, err)
	}
	ecomio.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) handleGetProductById(w http.ResponseWriter, r *http.Request) {
	// GetProductById(id int) (*Product, error)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ecomio.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	product, err := h.store.GetProductById(id)
	if err != nil {
		ecomio.WriteError(w, http.StatusBadRequest, err)
	}
	ecomio.WriteJSON(w, http.StatusOK, product)

}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.CreateProductPayload
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

	err := h.store.CreateProduct(types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
	})
	if err != nil {
		ecomio.WriteError(w, http.StatusBadRequest, err)
		return
	}
	ecomio.WriteJSON(w, http.StatusCreated, nil)
}
