package internalmodel

import (
	"time"

	"github.com/car-journal/lib/uuid"
)

const OauthAccessTokenTableName = "oauth_access_tokens"

type OauthAccessToken struct {
	ID        uuid.UUID  `json:"id" gorm:"primarykey"`
	UserID    uuid.UUID  `json:"user_id"`
	Revoked   bool       `json:"revoked"`
	CreatedAt time.Time  `json:"created_at" gorm:"<-:created"`
	UpdatedAt time.Time  `json:"updated_at"`
	ExpiresAt *time.Time `json:"expires_at"`
}
