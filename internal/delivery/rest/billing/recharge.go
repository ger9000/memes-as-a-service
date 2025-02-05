package billing

import (
	"encoding/json"
	"net/http"

	"github.com/ger9000/memes-as-a-service/internal/delivery/rest"
	router "github.com/ger9000/memes-as-a-service/internal/shared/http"
	"github.com/ger9000/memes-as-a-service/internal/usecases/billing"
)

type controller struct {
	action *billing.RechargeAvailableCallAction
}

func NewRechargeAvailableCallController(action *billing.RechargeAvailableCallAction) rest.IController {
	return &controller{
		action,
	}
}

func (c *controller) Invoke(w http.ResponseWriter, r *http.Request) {
	request := RechargeAvailableCallRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		router.BadRequest(w, r, err)
		return
	}

	if err := c.action.Do(r.Context(), request.Token, request.AmountToRecharge); err != nil {
		router.InternalServerError(w, r, nil)
		return
	}

	router.NoContent(w, r)
}
