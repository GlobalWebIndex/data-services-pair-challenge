package domain

import "context"

type AudienceRepository interface {
	Create(ctx context.Context, audience *Audience) error
	Update(ctx context.Context, audience *Audience) error
	Delete(ctx context.Context, audienceID AudienceID) error
	GetByID(ctx context.Context, audienceID AudienceID) (*Audience, error)
	GetAll(ctx context.Context) ([]*Audience, error)
}
