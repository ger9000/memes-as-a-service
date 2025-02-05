package rest

import (
	"net/http"

	"github.com/ger9000/memes-as-a-service/internal/shared/http/middleware"
	"github.com/go-chi/chi/v5"
)

type IController interface {
	Invoke(w http.ResponseWriter, r *http.Request)
}

type RouterHandlers struct {
	HealthCheckController           IController
	GetAllMemesController           IController
	RechargeAvailableCallController IController

	APICallsTrackerMiddleware middleware.APICallsTrackerMiddleware
}

func NewRouter(controllers RouterHandlers) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", controllers.HealthCheckController.Invoke)
	r.Post("/billing/recharge", controllers.RechargeAvailableCallController.Invoke)
	r.Route("/memes", func(r chi.Router) {
		r.Use(controllers.APICallsTrackerMiddleware.Validate)
		r.Get("/", controllers.GetAllMemesController.Invoke)
	})

	return r
}
