package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/neberson/pos-go-expert-fullcycle/tree/main/modulos/API/internal/dto"
	"github.com/neberson/pos-go-expert-fullcycle/tree/main/modulos/API/internal/entity"
	"github.com/neberson/pos-go-expert-fullcycle/tree/main/modulos/API/internal/infra/database"
	entityPkg "github.com/neberson/pos-go-expert-fullcycle/tree/main/modulos/API/pkg/entity"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create Product godoc
// @Summary Create a new product
// @Description Create a new product with the given details
// @Tags products
// @Accept json
// @Produce json
// @Param product body dto.CreateProductInput true "product request"
// @Success 201 {object} dto.CreateProductInput
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		errorResponse := Error{
			Message: "Invalid request body" + err.Error(),
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		errorResponse := Error{
			Message: "Invalid product data" + err.Error(),
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		errorResponse := Error{
			Message: "Failed to create product" + err.Error(),
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

// Get Product godoc
// @Summary Get a product by ID
// @Description Get a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID" Format(uuid)
// @Success 200 {object} entity.Product
// @Failure 204 {object} Error "No content"
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := Error{
			Message: "Product ID is required",
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	product, err := h.ProductDB.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		errorResponse := Error{
			Message: "Product not found",
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Update Product godoc
// @Summary Update product
// @Description Update product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID" Format(uuid)
// @Param request body dto.CreateProductInput true "product request"
// @Success 201 {object} dto.CreateProductInput
// @Failure 204 {object} Error "No content"
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{
			Message: "Product ID is required",
		}
		json.NewEncoder(w).Encode(error)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		error := Error{
			Message: "Invalid request body: " + err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}

	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		error := Error{
			Message: "Invalid product ID: " + err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, err = h.ProductDB.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		error := Error{
			Message: "Product not found",
		}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		error := Error{
			Message: "Failed to update product: " + err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Delete Product godoc
// @Summary Delete a product by ID
// @Description Delete a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID" Format(uuid)
// @Success 200
// @Success 204
// @Failure 400
// @Failure 500 {object} Error
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := h.ProductDB.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		errorResponse := Error{
			Message: "Failed to delete product: " + err.Error(),
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Get Product godoc
// @Summary get all products
// @Description get all products with pagination and sorting
// @Tags products
// @Accept json
// @Produce json
// @Param page query string false "Page number"
// @Param limit query string false "Limit per page"
// @Success 200 {array} entity.Product
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 {
		limitInt = 0
	}

	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := Error{
			Message: "Failed to retrieve products: " + err.Error(),
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if len(products) == 0 {
		w.WriteHeader(http.StatusNoContent)
		errorPResponse := Error{
			Message: "No products found",
		}
		json.NewEncoder(w).Encode(errorPResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
