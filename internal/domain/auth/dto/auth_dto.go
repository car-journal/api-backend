package authdto

import "time"

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type"`
}

type Me struct {
	ID                 string     `json:"id"`
	RoleID             string     `json:"role_id"`
	Email              string     `json:"email"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
	FirstName          string     `json:"first_name"`
	LastName           *string    `json:"last_name"`
	OauthAccessTokenID string     `json:"-"`
}
