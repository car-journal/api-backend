// Package database handles gorm
package database

import (
	"context"

	"gorm.io/gorm"
)

type dbKeyType struct{}

var dbKey dbKeyType

func Get(ctx context.Context) *gorm.DB {
	return ctx.Value(dbKey).(*gorm.DB)
}

func Set(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, dbKey, db)
}
