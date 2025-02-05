package tracker

import "context"

type IService interface {
	Find(context.Context, string) (*Track, error)
	ConsumeAvailableCall(context.Context, string) error
	Update(context.Context, *Track) error
}
