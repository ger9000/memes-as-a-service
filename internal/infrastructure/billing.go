package infrastructure

import (
	"context"

	"github.com/ger9000/memes-as-a-service/internal/domain/tracker"
	"gorm.io/gorm"
)

type BillingRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) tracker.IService {
	return &BillingRepository{db}
}

func (r *BillingRepository) Find(ctx context.Context, token string) (*tracker.Track, error) {
	track := tracker.Track{}
	res := r.db.First(&track, "token = ?", token)
	if res.Error != nil {
		if res.Error != gorm.ErrRecordNotFound {
			return nil, res.Error
		}
		return nil, nil
	}
	return &track, nil
}

func (r *BillingRepository) ConsumeAvailableCall(ctx context.Context, token string) error {
	track := tracker.Track{}
	res := r.db.Model(&track).
		Where("token = ?", token).
		UpdateColumn("available_calls", gorm.Expr("available_calls - 1"))
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *BillingRepository) Update(ctx context.Context, track *tracker.Track) error {
	res := r.db.Save(&track)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
