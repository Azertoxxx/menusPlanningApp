package router

import (
	"github.com/go-chi/chi/v5/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"

	"menusPlanningApp/api/controller"
)

type AppRouter struct {
	DishController     *controller.DishController
	PlanningController *controller.PlanningController
}

func New(app *AppRouter) *chi.Mux {
	r := chi.NewRouter()

	r.Use(corsMiddleware)
	r.Use(middleware.Logger)

	registerDishRoutes(r, app.DishController)
	registerPlanningRoutes(r, app.PlanningController)

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return r
}

func registerDishRoutes(r *chi.Mux, c *controller.DishController) {
	r.Route("/dishes", func(r chi.Router) {
		r.Get("/{id}", c.GetDish)
		r.Get("/", c.GetAllDishes)
		r.Post("/", c.CreateDish)
		r.Put("/{id}", c.UpdateDish)
		r.Delete("/{id}", c.DeleteDish)
	})
}

func registerPlanningRoutes(r *chi.Mux, c *controller.PlanningController) {
	r.Route("/plannings", func(r chi.Router) {
		r.Get("/{id}", c.GetPlanning)
		r.Get("/", c.GetAllPlannings)
		r.Post("/", c.CreatePlanning)
		r.Put("/{id}", c.UpdatePlanning)
		r.Delete("/{id}", c.DeletePlanning)
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
