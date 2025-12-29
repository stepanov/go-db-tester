package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	// clickhouse/sql driver may need an adapter; placeholder
	// _ "github.com/ClickHouse/clickhouse-go"

	"context"

	"github.com/go-redis/redis/v8"
)

// Manager holds connections to multiple databases
type Manager struct {
	Postgres   *sql.DB
	MySQL      *sql.DB
	Clickhouse *sql.DB // optionally use a dedicated client
	Redis      *redis.Client
}

// OpenAll opens all DB connections based on DSNs
func OpenAll(ctx context.Context, pgDSN, myDSN, chDSN, redisAddr string) (*Manager, error) {
	m := &Manager{}
	var err error
	if pgDSN != "" {
		m.Postgres, err = sql.Open("postgres", pgDSN)
		if err != nil {
			return nil, fmt.Errorf("open postgres: %w", err)
		}
	}
	if myDSN != "" {
		m.MySQL, err = sql.Open("mysql", myDSN)
		if err != nil {
			return nil, fmt.Errorf("open mysql: %w", err)
		}
	}
	if chDSN != "" {
		m.Clickhouse, err = sql.Open("clickhouse", chDSN)
		if err != nil {
			// driver placeholder - in real app use the clickhouse driver
			// for now just log missing driver error
			// return nil, fmt.Errorf("open clickhouse: %w", err)
		}
	}
	if redisAddr != "" {
		m.Redis = redis.NewClient(&redis.Options{Addr: redisAddr})
		if err := m.Redis.Ping(ctx).Err(); err != nil {
			return nil, fmt.Errorf("redis ping: %w", err)
		}
	}

	return m, nil
}

// Close closes all connections
func (m *Manager) Close() error {
	var err error
	if m.Postgres != nil {
		err = m.Postgres.Close()
	}
	if m.MySQL != nil {
		err = m.MySQL.Close()
	}
	if m.Clickhouse != nil {
		err = m.Clickhouse.Close()
	}
	if m.Redis != nil {
		err = m.Redis.Close()
	}
	return err
}
