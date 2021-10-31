package bootstrap

import (
	"context"
	"crudEmployee/config"
	"crudEmployee/internal/handler"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib" // driver sql
)

// Server struct http server
type Server struct {
	httpServer *http.Server
}

// Run http server
func (s *Server) Run(config *config.ServerConfig, handle *handler.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + config.Port,
		Handler:        handler.InitHandler(*handle),
		MaxHeaderBytes: config.MaxHeaderBytes << 20, // MB
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
	}
	return s.httpServer.ListenAndServe()
}

// Shutdown server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

// InitDB init postgres
func InitDB(config *config.DBConfig) (*sql.DB, error) {
	dbStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
	)
	db, err := sql.Open("pgx", dbStr)
	if err != nil {
		return nil, fmt.Errorf("connection error: %v ", err)
	}

	db.SetMaxOpenConns(config.MaxConn)
	db.SetMaxIdleConns(config.MaxIdleConn)
	db.SetConnMaxIdleTime(config.TimeIdleConn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("connection error: %v ", err)
	}

	return db, nil
}
