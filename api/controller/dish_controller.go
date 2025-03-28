package controller

import (
	"encoding/json"
	"menusPlanningApp/api/model"
	"menusPlanningApp/api/service"
	"menusPlanningApp/api/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type DishController struct {
	service *service.DishService
}

func NewDishController(service *service.DishService) *DishController {
	return &DishController{service}
}

func (c *DishController) GetAllDishes(w http.ResponseWriter, r *http.Request) {
	errors := []string{}

	limit, err := utils.GetQueryInt(r, "limit", 10)
	if err != nil {
		errors = append(errors, "Invalid 'limit': must be a positive integer")
	}
	offset, err := utils.GetQueryInt(r, "offset", 0)
	if err != nil {
		errors = append(errors, "Invalid 'offset': must be a non-negative integer")
	}
	sortOrder := utils.GetQueryString(r, "sortOrder", "asc")
	if sortOrder != "asc" && sortOrder != "desc" {
		errors = append(errors, "Invalid 'sortOrder': must be 'asc' or 'desc'")
	}
	if len(errors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":  "Invalid query parameters",
			"issues": errors,
		})
		return
	}

	filters := utils.GetQueryFilters(r)
	queryParams := model.QueryParams{
		Limit:     limit,
		Offset:    offset,
		SortBy:    utils.GetQueryString(r, "sortBy", "id"),
		SortOrder: sortOrder,
		Filters:   filters,
	}
	dishes, err := c.service.FindAll(queryParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dishes)
}

func (c *DishController) CreateDish(w http.ResponseWriter, r *http.Request) {
	var dish model.Dish
	if err := json.NewDecoder(r.Body).Decode(&dish); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdDish, err := c.service.Create(&dish)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdDish)
}

func (c *DishController) GetDish(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}
	dish, err := c.service.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dish)
}

func (c *DishController) UpdateDish(w http.ResponseWriter, r *http.Request) {
	var dish model.Dish
	if err := json.NewDecoder(r.Body).Decode(&dish); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedDish, err := c.service.Update(&dish)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedDish)
}

func (c *DishController) DeleteDish(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}
	if err := c.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
