package database

import (
	"context"
	"counter/config"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB(cfg *config.Config) *PostgresDB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresSSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("failed to ping PostgreSQL: %v", err)
	}

	log.Println("Connected to PostgreSQL")

	return &PostgresDB{
		DB: db,
	}
}

func (db *PostgresDB) Migrate(cfg *config.Config) {
	migrateURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB, cfg.PostgresSSLMode)

	m, err := migrate.New("file://./migrations", migrateURL)
	if err != nil {
		log.Fatalf("could not create migration instance: %v", err)
	}

	if err = m.Up(); err != nil {
		log.Printf("could not apply migrations: %v", err)
	}
}
