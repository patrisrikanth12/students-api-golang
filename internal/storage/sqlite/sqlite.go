package sqlite

import (
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
	"github.com/patrisrikanth12/students-api-golang/internal/config"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)

	if err != nil {
		return nil, err 
	}

	slog.Info("Initialized DB...")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS students (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			email TEXT,
			mobile TEXT
		)
	`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{Db: db},nil
}

func (s *Sqlite) CreateStudent(name string, email string, mobile string) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO students(name, email, mobile) Values(?, ?, ?)")
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(name, email, mobile)
	if err != nil {
		return 0, err 
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err 
	}


	return lastId, nil
}