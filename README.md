<<<<<<< HEAD
# go-clickhouse (scaffold)

This is a scaffold for a Go service that can connect to multiple databases and message brokers (Postgres, MySQL, ClickHouse, Redis, Kafka, NATS, RabbitMQ).

Structure
- `cmd/server` - application entrypoint
- `internal/config` - configuration loader (Viper)
- `internal/db` - multi-db connection manager
- `internal/broker` - broker interfaces and implementations
- `migrations` - DB migrations
- `scripts` - helper scripts

Quickstart
1. Copy `.env.example` to `.env` and adjust DSNs.
2. Start dependencies: `docker-compose up -d`
3. Run the server: `go run ./cmd/server`

This scaffold provides placeholder implementations for brokers and clickhouse which you should replace with real clients (sarama/confluent, nats.go, amqp, clickhouse-go, etc.).
=======
# go-db-tester
>>>>>>> upstream/master
