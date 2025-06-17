package database

import (
	"context"

	"github.com/car-journal/config"
)

func Run(ctx context.Context, fn func(ctx context.Context) error) error {
	db := config.ConnectGormPostgres()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	ctx = Set(ctx, tx)
	if err := fn(ctx); nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
