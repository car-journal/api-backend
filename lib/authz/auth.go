package authz

import (
	"context"
	"time"
)

type authzKeyType struct{}

var authzKey authzKeyType

type Authz struct {
	ID                 string     `json:"id"`
	Email              string     `json:"email"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
	FirstName          string     `json:"first_name"`
	LastName           *string    `json:"last_name"`
	PictureURL         *string    `json:"picture_url"`
	OauthAccessTokenID string     `json:"-"`
}

func GetAuthUser(ctx context.Context) Authz {
	return ctx.Value(authzKey).(Authz)
}

func SetAuthUser(ctx context.Context, user Authz) context.Context {
	return context.WithValue(ctx, authzKey, user)
}
