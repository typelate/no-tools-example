package main

import (
	"cmp"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/typelate/no-tools-example/internal/database"
	"github.com/typelate/no-tools-example/internal/hypertext"
)

func main() {
	startupCTX, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	conn, err := pgx.Connect(startupCTX, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer closeWithContextAndIgnoreError(context.Background(), conn)
	if err := conn.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}
	if err := database.Migrate(*conn.Config()); err != nil {
		log.Fatal(err)
	}
	srv := &hypertext.Server{
		DBQuery: database.New(),
		DBConn:  conn,
	}
	mux := http.NewServeMux()
	hypertext.TemplateRoutes(mux, srv)
	log.Fatal(http.ListenAndServe(":"+cmp.Or(os.Getenv("PORT"), "8080"), mux))
}

func closeWithContextAndIgnoreError(ctx context.Context, c interface {
	Close(ctx context.Context) error
}) {
	_ = c.Close(ctx)
}
