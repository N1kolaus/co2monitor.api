package main

import (
	"context"
	"database/sql"
	"expvar"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/FMinister/co2monitor-api/internal/data"
	"github.com/FMinister/co2monitor-api/internal/vcs"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	version = vcs.Version()
)

type config struct {
	port string
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTIme  time.Duration
	}
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
	cors struct {
		trustedOrigins []string
	}
	maxTimeFrameDays int
}

type application struct {
	config config
	logger *slog.Logger
	models data.Models
	wg     sync.WaitGroup
}

func main() {
	godotenv.Load()
	var cfg config

	flag.StringVar(&cfg.port, "port", os.Getenv("PORT"), "API server port")
	flag.StringVar(&cfg.env, "env", os.Getenv("APP_ENV"), "Environment (development|staging|production)")

	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("CO2MONITOR_DATABASE_DSN"), "PostgreSQL DSN")

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(&cfg.db.maxIdleTIme, "db-max-idle-time", 15*time.Minute, "PostgreSQL max connections idle time")

	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")

	flag.IntVar(&cfg.maxTimeFrameDays, "maxTimeFrameDays", 14, "max time frame days for co2 data query")

	flag.Func("cors-trusted-origins", "Trusted CORS origins (space seperated)", func(val string) error {
		cfg.cors.trustedOrigins = strings.Fields(val)
		return nil
	})

	displayVersion := flag.Bool("version", false, "Display version and exit")

	flag.Parse()

	// If the version flag value is true, then print out the version number and immediately exit.
	if *displayVersion {
		fmt.Printf("Version:\t%s\n", version)
		os.Exit(0)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	db, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("database connection pool established")

	expvar.NewString("version").Set(version)
	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))
	expvar.Publish("database", expvar.Func(func() any {
		return db.Stats()
	}))
	expvar.Publish("timestamp", expvar.Func(func() any {
		return time.Now().Unix()
	}))

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	err = app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	db.SetConnMaxIdleTime(cfg.db.maxIdleTIme)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
