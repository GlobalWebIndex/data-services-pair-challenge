package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/GlobalWebIndex/data-services-pair-challenge/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AudienceRepository struct {
	db *pgxpool.Pool
}

func NewAudienceRepository(db *pgxpool.Pool) *AudienceRepository {
	return &AudienceRepository{
		db: db,
	}
}

func (r *AudienceRepository) Create(ctx context.Context, audience *domain.Audience) error {
	// Implementation to insert the Audience into the database.
	// Use prepared statements to prevent SQL injection.

	query := `
		INSERT INTO audiences (name, expression)
		VALUES ($1, $2)
	`

	_, err := r.db.Exec(ctx, query, audience.Name, audience.Expression)
	if err != nil {
		return fmt.Errorf("failed to create Audience: %w", err)
	}

	return nil
}

func (r *AudienceRepository) Update(ctx context.Context, audience *domain.Audience) error {
	// Implementation to update the Audience in the database.
	// Use prepared statements to prevent SQL injection.

	query := `
		UPDATE audiences
		SET name = $2, expression = $3
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query, audience.ID, audience.Name, audience.Expression)
	if err != nil {
		return fmt.Errorf("failed to update Audience: %w", err)
	}

	return nil
}

func (r *AudienceRepository) Delete(ctx context.Context, audienceID domain.AudienceID) error {
	// Implementation to delete the Audience from the database.
	// Use prepared statements to prevent SQL injection.

	query := `
		DELETE FROM audiences
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query, audienceID)
	if err != nil {
		return fmt.Errorf("failed to delete Audience: %w", err)
	}

	return nil
}

func (r *AudienceRepository) GetByID(ctx context.Context, audienceID domain.AudienceID) (*domain.Audience, error) {
	// Implementation to get an Audience by its ID from the database.
	// Use prepared statements to prevent SQL injection.

	query := `
		SELECT id, name, expression
		FROM audiences
		WHERE id = $1
	`

	row := r.db.QueryRow(ctx, query, audienceID)
	audience := &domain.Audience{}
	err := row.Scan(&audience.ID, &audience.Name, &audience.Expression)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("audience not found")
		}
		return nil, fmt.Errorf("failed to get Audience by ID: %w", err)
	}

	return audience, nil
}

func (r *AudienceRepository) GetAll(ctx context.Context) ([]*domain.Audience, error) {
	// Implementation to get all Audiences from the database.
	// Use prepared statements to prevent SQL injection.

	query := `
		SELECT id, name, expression
		FROM audiences
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all Audiences: %w", err)
	}
	defer rows.Close()

	audiences := []*domain.Audience{}
	for rows.Next() {
		audience := &domain.Audience{}
		err := rows.Scan(&audience.ID, &audience.Name, &audience.Expression)
		if err != nil {
			return nil, fmt.Errorf("failed to scan Audience: %w", err)
		}
		audiences = append(audiences, audience)
	}

	return audiences, nil
}
