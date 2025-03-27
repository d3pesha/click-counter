package storage

import (
	"context"
	"counter/internal/model"
	"database/sql"
	"errors"
	"fmt"
	"time"
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

func (r *bannerStorage) IncrementClick(ctx context.Context, bannerID int, timestamp time.Time, clickCount int) error {
	query := `INSERT INTO banner_clicks (banner_id, timestamp, click_count) 
              VALUES ($1, $2, $3)
              ON CONFLICT (banner_id, timestamp) 
              DO UPDATE SET click_count = banner_clicks.click_count + EXCLUDED.click_count`

	_, err := r.db.ExecContext(ctx, query, bannerID, timestamp, clickCount)
	if err != nil {
		return fmt.Errorf("failed to increment click: %w", err)
	}

	return nil
}

func (r *bannerStorage) GetStats(ctx context.Context, bannerID int, from, to time.Time) ([]*model.BannerClick, error) {
	query := `SELECT banner_id, timestamp, click_count FROM banner_clicks 
              WHERE banner_id = $1 AND timestamp BETWEEN $2 AND $3`

	rows, err := r.db.QueryContext(ctx, query, bannerID, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var rowCount int
	err = r.db.QueryRowContext(ctx, `
        SELECT COUNT(*)
        FROM banner_clicks
        WHERE banner_id = $1 AND timestamp BETWEEN $2 AND $3
    `, bannerID, from, to).Scan(&rowCount)
	if err != nil {
		return nil, fmt.Errorf("failed to count rows: %w", err)
	}

	stats := make([]*model.BannerClick, 0, rowCount)
	for rows.Next() {
		var click model.BannerClick
		if err = rows.Scan(&click.BannerID, &click.Timestamp, &click.ClickCount); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		stats = append(stats, &click)
	}

	return stats, nil
}
