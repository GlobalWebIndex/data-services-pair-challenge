package domain

import (
	"context"
	"fmt"
)

type AudienceService struct {
	audienceRepository AudienceRepository
}

func NewAudienceService(audienceRepository AudienceRepository) *AudienceService {
	return &AudienceService{
		audienceRepository: audienceRepository,
	}
}

// CreateAudience creates a new Audience after performing basic validations.
func (s *AudienceService) CreateAudience(ctx context.Context, name AudienceName, expression Expression) (*Audience, error) {
	// Perform basic validation before creating the Audience.
	if name == "" {
		return nil, fmt.Errorf("Audience name cannot be empty")
	}

	// Create the Audience entity.
	audience := Audience{
		Name:       name,
		Expression: expression,
	}

	// Now, call the repository's Create method to persist the Audience in the database.
	err := s.audienceRepository.Create(ctx, &audience)
	if err != nil {
		// Handle any error occurred during the repository call.
		return nil, err
	}

	return &audience, nil
}

// GetAudienceByID fetches an audience by its ID.
func (s *AudienceService) GetAudienceByID(ctx context.Context, audienceID AudienceID) (*Audience, error) {
	return s.audienceRepository.GetByID(ctx, audienceID)
}
