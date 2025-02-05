package billing

import (
	"context"
	"errors"

	"github.com/ger9000/memes-as-a-service/internal/domain/tracker"
	"github.com/rs/zerolog/log"
)

type ConsumeAvailableCallAction struct {
	service tracker.IService
}

func NewConsumeAvailableCallAction(service tracker.IService) *ConsumeAvailableCallAction {
	return &ConsumeAvailableCallAction{service}
}

func (a *ConsumeAvailableCallAction) Do(ctx context.Context, token string) error {
	if err := a.service.ConsumeAvailableCall(ctx, token); err != nil {
		log.Error().Err(err).Send()
		return errors.New("error saving api call tracking")
	}

	return nil
}
