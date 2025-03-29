package storage

import (
	"context"
	"counter/internal/model"
	"database/sql"
	"errors"
	"fmt"
)

type bannerStorage struct {
	db *sql.DB
}

func NewBannerStorage(db *sql.DB) BannerStorage {
	return &bannerStorage{db: db}
}

func (r *bannerStorage) GetByID(ctx context.Context, bannerID int) (*model.Banner, error) {
	query := `SELECT id, name FROM banners WHERE id = $1`

	var banner model.Banner
	err := r.db.QueryRowContext(ctx, query, bannerID).Scan(&banner.ID, &banner.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("banner with ID %d not found", bannerID)
		}

		return nil, fmt.Errorf("failed to get banner by ID: %w", err)
	}

	return &banner, err
}

func (r *bannerStorage) GetAllIDs(ctx context.Context) ([]int, error) {
	query := `SELECT id FROM banners`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get banner IDs: %w", err)
	}
	defer rows.Close()

	var count int
	countQuery := `SELECT COUNT(*) FROM banners`
	err = r.db.QueryRowContext(ctx, countQuery).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("failed to count banners: %w", err)
	}

	ids := make([]int, 0, count)
	for rows.Next() {
		var id int
		if err = rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to scan banner ID: %w", err)
		}

		ids = append(ids, id)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return ids, nil
}
