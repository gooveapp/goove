package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

type DB struct {
	*sql.DB
}

func New(dbPath string) (*DB, error) {
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	log.Println("Database connection established")

	return &DB{db}, nil
}

func (db *DB) Close() error {
	return db.DB.Close()
}

// RunMigrations runs all database migrations
func (db *DB) RunMigrations() error {
	// Create migrations table if it doesn't exist
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version INTEGER PRIMARY KEY,
			applied_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	// Get list of applied migrations
	applied := make(map[int]bool)
	rows, err := db.Query("SELECT version FROM schema_migrations")
	if err != nil {
		return fmt.Errorf("failed to query migrations: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var version int
		if err := rows.Scan(&version); err != nil {
			return fmt.Errorf("failed to scan migration version: %w", err)
		}
		applied[version] = true
	}

	// Define migrations
	migrations := []struct {
		version int
		name    string
		sql     string
	}{
		{1, "create_artists", `
			CREATE TABLE IF NOT EXISTS artists (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT NOT NULL UNIQUE,
				sort_name TEXT,
				discogs_id INTEGER,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
			);
			CREATE INDEX IF NOT EXISTS idx_artists_name ON artists(name);
		`},
		{2, "create_labels", `
			CREATE TABLE IF NOT EXISTS labels (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT NOT NULL UNIQUE,
				discogs_id INTEGER,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
			);
			CREATE INDEX IF NOT EXISTS idx_labels_name ON labels(name);
		`},
		{3, "create_genres", `
			CREATE TABLE IF NOT EXISTS genres (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT NOT NULL UNIQUE,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP
			);
			CREATE INDEX IF NOT EXISTS idx_genres_name ON genres(name);
		`},
		{4, "create_records", `
			CREATE TABLE IF NOT EXISTS records (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				title TEXT NOT NULL,
				artist_id INTEGER NOT NULL,
				label_id INTEGER,
				catalog_number TEXT,
				barcode TEXT,
				release_year INTEGER,
				country TEXT,
				format TEXT,
				sleeve_condition TEXT,
				media_condition TEXT,
				notes TEXT,
				purchase_date DATE,
				purchase_price DECIMAL(10,2),
				purchase_location TEXT,
				cover_image_path TEXT,
				discogs_release_id INTEGER,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (artist_id) REFERENCES artists(id) ON DELETE RESTRICT,
				FOREIGN KEY (label_id) REFERENCES labels(id) ON DELETE SET NULL
			);
			CREATE INDEX IF NOT EXISTS idx_records_artist ON records(artist_id);
			CREATE INDEX IF NOT EXISTS idx_records_label ON records(label_id);
			CREATE INDEX IF NOT EXISTS idx_records_title ON records(title);
			CREATE INDEX IF NOT EXISTS idx_records_barcode ON records(barcode);
		`},
		{5, "create_record_genres", `
			CREATE TABLE IF NOT EXISTS record_genres (
				record_id INTEGER NOT NULL,
				genre_id INTEGER NOT NULL,
				PRIMARY KEY (record_id, genre_id),
				FOREIGN KEY (record_id) REFERENCES records(id) ON DELETE CASCADE,
				FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE
			);
			CREATE INDEX IF NOT EXISTS idx_record_genres_record ON record_genres(record_id);
			CREATE INDEX IF NOT EXISTS idx_record_genres_genre ON record_genres(genre_id);
		`},
		{6, "create_user_settings", `
			CREATE TABLE IF NOT EXISTS user_settings (
				key TEXT PRIMARY KEY,
				value TEXT NOT NULL,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
			);
			INSERT OR IGNORE INTO user_settings (key, value) VALUES 
				('currency', 'GBP'),
				('date_format', 'DD/MM/YYYY'),
				('default_condition', 'VG+');
		`},
	}

	// Run migrations in order
	for _, m := range migrations {
		if applied[m.version] {
			continue
		}

		log.Printf("Running migration %d: %s", m.version, m.name)

		// Start transaction
		tx, err := db.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}

		// Run migration
		if _, err := tx.Exec(m.sql); err != nil {
			tx.Rollback()
			return fmt.Errorf("migration %d failed: %w", m.version, err)
		}

		// Record migration
		if _, err := tx.Exec("INSERT INTO schema_migrations (version) VALUES (?)", m.version); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to record migration: %w", err)
		}

		// Commit transaction
		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit migration: %w", err)
		}

		log.Printf("Migration %d completed: %s", m.version, m.name)
	}

	log.Println("All migrations completed")
	return nil
}
