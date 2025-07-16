package database

import (
	"context"
)

func Run(ctx context.Context, fn func(ctx context.Context) error) error {
	db := Get(ctx)
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
