package database

import (
	"context"
	"database/sql"
	"fmt"

	"identity-service/config"

	_ "github.com/lib/pq"
)

// Database wraps the sql.DB with additional functionality
type Database struct {
	DB     *sql.DB
	config *config.DatabaseConfig
}

// NewDatabase creates a new database connection with proper configuration
func NewDatabase(cfg *config.DatabaseConfig) (*Database, error) {
	db, err := sql.Open("postgres", cfg.ConnectionString())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Database{
		DB:     db,
		config: cfg,
	}, nil
}

// Close closes the database connection
func (d *Database) Close() error {
	return d.DB.Close()
}

// InitSchema initializes the database schema with transaction support
func (d *Database) InitSchema(ctx context.Context) error {
	// Use transaction for schema initialization
	tx, err := d.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback() // Will rollback if not committed

	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) NOT NULL UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	if _, err := tx.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	// Create indexes for better performance
	indexQuery := `
		CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
	`
	if _, err := tx.ExecContext(ctx, indexQuery); err != nil {
		return fmt.Errorf("failed to create email index: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit schema transaction: %w", err)
	}

	return nil
}

// BeginTx starts a new transaction
func (d *Database) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return d.DB.BeginTx(ctx, opts)
}

// ExecContext executes a query without returning any rows
func (d *Database) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.DB.ExecContext(ctx, query, args...)
}

// QueryContext executes a query that returns rows
func (d *Database) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return d.DB.QueryContext(ctx, query, args...)
}

// QueryRowContext executes a query that returns at most one row
func (d *Database) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return d.DB.QueryRowContext(ctx, query, args...)
}
