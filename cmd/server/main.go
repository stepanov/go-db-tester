package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-databases-test/internal/broker"
	"go-databases-test/internal/config"
	"go-databases-test/internal/db"
	"go-databases-test/internal/log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("load config:", err)
	}

	ctx := context.Background()
	mgr, err := db.OpenAll(ctx, cfg.PostgresDSN, cfg.MySQLDSN, cfg.Clickhouse, cfg.RedisAddr)
	if err != nil {
		log.Fatal("open dbs:", err)
	}
	defer mgr.Close()

	providers := &broker.Provider{}
	if len(cfg.KafkaBrokers) > 0 {
		k, _ := broker.NewKafka(cfg.KafkaBrokers)
		providers.Kafka = k
	}
	if cfg.NatsURL != "" {
		n, _ := broker.NewNats(cfg.NatsURL)
		providers.Nats = n
	}
	if cfg.RabbitURL != "" {
		r, _ := broker.NewRabbit(cfg.RabbitURL)
		providers.Rabbit = r
	}

	// Minimal HTTP server
	srv := &http.Server{
		Addr:         cfg.HTTP.Addr,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
		Handler:      http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "ok") }),
	}

	go func() {
		log.Info("server starting on", cfg.HTTP.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("server error:", err)
		}
	}()

	// Wait for signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("shutting down")
	ctxShut, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctxShut)
	_ = providersClose(providers)
	log.Info("done")
}

func providersClose(p *broker.Provider) error {
	if p.Kafka != nil {
		_ = p.Kafka.Close()
	}
	if p.Nats != nil {
		_ = p.Nats.Close()
	}
	if p.Rabbit != nil {
		_ = p.Rabbit.Close()
	}
	return nil
}
