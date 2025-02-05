package billing

import (
	"context"
	"errors"

	"github.com/ger9000/memes-as-a-service/internal/domain/tracker"
	"github.com/rs/zerolog/log"
)

type RechargeAvailableCallAction struct {
	service tracker.IService
}

func NewRechargeAvailableCallAction(service tracker.IService) *RechargeAvailableCallAction {
	return &RechargeAvailableCallAction{service}
}

func (a *RechargeAvailableCallAction) Do(ctx context.Context, token string, amountToRecharge int32) error {
	track, err := a.service.Find(context.Background(), token)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	if track == nil {
		track = &tracker.Track{
			Token: token,
		}
	}

	track.AvailableCalls += amountToRecharge

	if err := a.service.Update(ctx, track); err != nil {
		log.Error().Err(err).Send()
		return errors.New("error saving api call tracking")
	}

	return nil
}
