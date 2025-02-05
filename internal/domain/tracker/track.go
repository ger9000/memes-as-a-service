package tracker

import (
	"time"

	"gorm.io/gorm"
)

type Track struct {
	Token          string    `gorm:"primaryKey;column:token"`
	AvailableCalls int32     `gorm:"column:available_calls"`
	LastRecharge   time.Time `gorm:"column:last_recharge"`
}

func (b *Track) BeforeUpdate(tx *gorm.DB) error {
	b.LastRecharge = time.Now()

	return nil
}
