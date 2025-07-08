// Package authrepository handles db queries for auth domain
package authrepository

import (
	"context"

	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/database"
)

type Interface interface {
	FindActiveOauthAccessTokenByID(ctx context.Context, id string) (*internalmodel.OauthAccessToken, error)
	CreateOauthAccessToken(ctx context.Context, oauthAccessToken *internalmodel.OauthAccessToken) (*internalmodel.OauthAccessToken, error)
	RevokeOauthAccessTokensByUserID(ctx context.Context, userID string) error
}

type repository struct{}

func Repository() Interface {
	return &repository{}
}

func (r repository) FindActiveOauthAccessTokenByID(ctx context.Context, id string) (*internalmodel.OauthAccessToken, error) {
	var result *internalmodel.OauthAccessToken
	db := database.Get(ctx)
	db = database.EqualsTo(db, "id", id)
	db = database.EqualsTo(db, "revoked", false)
	db = database.GreaterThan(db, "expires_at", "NOW()")

	if err := db.First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r repository) CreateOauthAccessToken(ctx context.Context, oauthAccessToken *internalmodel.OauthAccessToken) (*internalmodel.OauthAccessToken, error) {
	db := database.Get(ctx)
	if err := db.Create(&oauthAccessToken).Error; err != nil {
		return nil, err
	}
	return oauthAccessToken, nil
}

func (r repository) RevokeOauthAccessTokensByUserID(ctx context.Context, userID string) error {
	db := database.Get(ctx)
	db = database.EqualsTo(db, "user_id", userID)
	db = database.EqualsTo(db, "revoked", false)
	return db.Model(&internalmodel.OauthAccessToken{}).Update("Revoked", true).Error
}
