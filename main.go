package main

import (
	"context"
	"csms/backend/db"
	"csms/backend/db/sqlc"
	"csms/backend/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	l := log.New(os.Stdout, "CSMS-SERVER-", log.LstdFlags)
	ctx := context.Background()

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	// load env
	err := godotenv.Load(env + ".env")
	if err != nil {
		l.Fatalf("Error reding the .env %s", err)
	}

	// CRUD
	db, queries := db.Instantiate(l)
	if db == nil || queries == nil {
		l.Println("Exiting due to database connection error")
		return
	}
	defer db.Close()

	server := &http.Server{
		Addr:        "0.0.0.0:" + os.Getenv("PORT"),
		Handler:     defineMultiplexer(l, queries),
		IdleTimeout: 30 * time.Second,
		ReadTimeout: time.Second,
	}

	go startServer(server, l)

	shut := make(chan os.Signal, 1)
	signal.Notify(shut, syscall.SIGINT, syscall.SIGTERM)

	<-shut

	timeout_ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stopServer(server, l, &timeout_ctx, &cancel)
}

func startServer(s *http.Server, l *log.Logger) {
	l.Println("🔥 Server is running on", s.Addr)

	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		l.Fatalln("Server is failed due to", err)
	}
}

func stopServer(s *http.Server, l *log.Logger, ctx *context.Context, cancel *context.CancelFunc) {
	l.Println("💅 Shutting down the server")
	s.Shutdown(*ctx)
	c := *cancel
	c()
}

func defineMultiplexer(l *log.Logger, q *sqlc.Queries) http.Handler {

	// reference to the handler
	hello_handler := handlers.NewHello(l)

	// handle multiplexer
	mux := http.NewServeMux()
	mux.Handle("/", hello_handler)

	corsMiddleware := cors.AllowAll().Handler

	handler := corsMiddleware(mux)

	return handler
}
