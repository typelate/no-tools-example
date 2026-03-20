package database

import (
	"context"
	"embed"
	"io/fs"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/typelate/loosey"
)

//go:embed migrations/*.sql
var migrations embed.FS

func Migrate(ctx context.Context, config pgx.ConnConfig) error {
	dir, err := fs.Sub(migrations, "migrations")
	if err != nil {
		return err
	}
	db := stdlib.OpenDB(config)
	defer func() {
		_ = db.Close()
	}()
	m, err := loosey.NewPostgres(ctx, db, dir)
	if err != nil {
		return err
	}
	_, err = m.Up(ctx)
	if err != nil {
		return err
	}
	return nil
}
