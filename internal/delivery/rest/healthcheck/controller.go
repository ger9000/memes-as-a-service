package healthcheck

import (
	"net/http"

	"github.com/ger9000/memes-as-a-service/internal/delivery/rest"
	"github.com/go-chi/render"
)

type controller struct{}

func (*controller) Invoke(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]string{
		"env":  "DEV",
		"name": "MaaS (memes as a service)",
	})
}

func NewController() rest.IController {
	return new(controller)
}
