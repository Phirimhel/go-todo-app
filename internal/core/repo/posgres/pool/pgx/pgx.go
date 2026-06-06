package core_pgx_pool

import (
	"context"
	"fmt"
	"time"

	core_postgres_pool "github.com/Phirimhel/go-todo-app/internal/core/repo/posgres/pool"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxConnectionPool struct {
	*pgxpool.Pool
	opTimeout time.Duration
}

func NewPgxConnectionPool(ctx context.Context, conf Config) (*PgxConnectionPool, error) {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)

	pgxconfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("parse pgx config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxconfig)
	if err != nil {
		return nil, fmt.Errorf("create pgx pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pgx pool ping: %w", err)
	}

	return &PgxConnectionPool{
		Pool:      pool,
		opTimeout: conf.Timeout,
	}, nil

}

func (c *PgxConnectionPool) Query(ctx context.Context, sql string, args ...any) (core_postgres_pool.Rows, error) {
	rows, err := c.Pool.Query(ctx, sql, args...)
	if err != nil {
		return pgxRows{}, err
	}
	return pgxRows{rows}, nil
}

func (c *PgxConnectionPool) QueryRow(ctx context.Context, sql string, args ...any) core_postgres_pool.Row {
	row := c.Pool.QueryRow(ctx, sql, args...)
	return pgxRow{row}
}

func (c *PgxConnectionPool) Exec(ctx context.Context, sql string, args ...any) (core_postgres_pool.CommandTag, error) {
	cmdTag, err := c.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return pgxCommandTag{}, err
	}

	return pgxCommandTag{cmdTag}, nil
}

func (c *PgxConnectionPool) Close() {
	c.Pool.Close()
}

func (c *PgxConnectionPool) OpTimeout() time.Duration {
	return c.opTimeout
}
