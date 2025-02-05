package middleware

import (
	"errors"
	"net/http"

	router "github.com/ger9000/memes-as-a-service/internal/shared/http"
	"github.com/ger9000/memes-as-a-service/internal/usecases/billing"
)

type APICallsTrackerMiddleware struct {
	findAction    *billing.FindTrackAction
	consumeAction *billing.ConsumeAvailableCallAction
}

func NewCallsTracker(findAction *billing.FindTrackAction, consumeAction *billing.ConsumeAvailableCallAction) *APICallsTrackerMiddleware {
	return &APICallsTrackerMiddleware{findAction, consumeAction}
}

func (m *APICallsTrackerMiddleware) Validate(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		headers := r.Header
		if len(headers["Authorization"]) > 0 {
			token := headers["Authorization"][0]
			track, err := m.findAction.Do(r.Context(), token)
			if err != nil {
				router.InternalServerError(w, r, nil)
				return
			}

			if track == nil || track.AvailableCalls <= 0 {
				router.BadRequest(w, r, errors.New("no available api calls"))
				return
			}

			if err := m.consumeAction.Do(r.Context(), track.Token); err != nil {
				router.InternalServerError(w, r, nil)
				return
			}
		} else {
			router.Unauthorized(w, r, errors.New("missing authorization token"))
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
