package billing

import (
	"context"
	"errors"

	"github.com/ger9000/memes-as-a-service/internal/domain/tracker"
	"github.com/rs/zerolog/log"
)

type FindTrackAction struct {
	service tracker.IService
}

func NewFindTrackAction(service tracker.IService) *FindTrackAction {
	return &FindTrackAction{service}
}

func (a *FindTrackAction) Do(ctx context.Context, token string) (*tracker.Track, error) {
	track, err := a.service.Find(ctx, token)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, errors.New("error saving api call tracking")
	}

	return track, nil
}
