package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

type Postgres struct {
	maxPoolSize int
	connAtempts int
	conTimeout  time.Duration
}

func New(config string) (*pgx.Conn, error) {
	//TODO pgxpool
	conn, err := pgx.Connect(context.Background(), config)
	if err != nil {
		return conn, fmt.Errorf("postgres init error: %w", err)
	}
	return conn, nil
}
