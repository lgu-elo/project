package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/lgu-elo/project/internal/config"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	defaultAttempts = 10
	defaultTimeout  = time.Second

	Up   = operation("up")
	Down = operation("down")
)

type operation string

func main() {
	cfg := config.New()
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)

	Do(Up, "./migrations", dsn)
}

func Do(op operation, dir, connString string) {
	m, err := connect(dir, connString)
	if err != nil {
		log.Fatalf("migrate: postgres connect error: %s", err.Error())
	}
	defer m.Close()

	if op == Up {
		err = m.Up()
	}

	if op == Down {
		err = m.Up()
	}

	if errors.Is(err, migrate.ErrNoChange) {
		m.Close()
		log.Printf("migrate: %s no change", op)
	}

	if err != nil {
		m.Close()
		log.Fatalf("migrate: %s error: %s", op, err.Error())
	}

	m.Close()
}

func connect(dir, connString string) (*migrate.Migrate, error) {
	var (
		attempts = defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://"+dir, connString)
		if err != nil {
			return nil, err
		}

		if err == nil {
			break
		}

		log.Printf("migrate: trying connecting to postgres, attempts left: %d", attempts)
		time.Sleep(defaultTimeout)
		attempts--
	}

	if err != nil {
		return nil, err
	}

	return m, nil
}
