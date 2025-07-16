package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client struct {
	conn    *pgx.Conn
	Queries *Queries
	Pool    *pgxpool.Pool
}

func Connect(ctx context.Context, cfg Config) (dbClient *Client, fault error) {
	c, err := pgx.Connect(ctx, cfg.GetURI())
	if err != nil {
		return nil, fmt.Errorf("could not connect to postgres: %w", err)
	}

	p, err := pgxpool.New(ctx, cfg.GetURI())
	if err != nil {
		return nil, fmt.Errorf("could not create connection pool: %w", err)
	}

	q := New(c)

	return &Client{
		conn:    c,
		Queries: q,
		Pool:    p,
	}, nil
}

func (c *Client) Close() {
	c.conn.Close(context.Background())
	c.Pool.Close()
}

func (c *Client) GetQueries() *Queries {
	return c.Queries
}

func (c *Client) GetPool() *pgxpool.Pool {
	return c.Pool
}
