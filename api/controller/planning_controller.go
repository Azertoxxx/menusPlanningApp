package controller

import (
	"encoding/json"
	"menusPlanningApp/api/model"
	"menusPlanningApp/api/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type PlanningController struct {
	service *service.PlanningService
}

func NewPlanningController(service *service.PlanningService) *PlanningController {
	return &PlanningController{service}
}

func (c *PlanningController) GetAllPlannings(w http.ResponseWriter, r *http.Request) {
	plannings, err := c.service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(plannings)
}

func (c *PlanningController) CreatePlanning(w http.ResponseWriter, r *http.Request) {
	var planning model.Planning
	if err := json.NewDecoder(r.Body).Decode(&planning); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdPlanning, err := c.service.Create(&planning)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPlanning)
}

func (c *PlanningController) GetPlanning(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}
	planning, err := c.service.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(planning)
}

func (c *PlanningController) UpdatePlanning(w http.ResponseWriter, r *http.Request) {
	var planning model.Planning
	if err := json.NewDecoder(r.Body).Decode(&planning); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedPlanning, err := c.service.Update(&planning)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedPlanning)
}

func (c *PlanningController) DeletePlanning(w http.ResponseWriter, r *http.Request) {
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
