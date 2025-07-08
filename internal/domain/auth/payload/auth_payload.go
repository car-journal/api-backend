// Package authpayload handles payload for auth domain
package authpayload

import "time"

type RegisterPayload struct {
	ClientSecret string `json:"client_secret" validate:"required"`

	Email           string     `json:"email" validate:"required"`
	Password        string     `json:"password" validate:"required,min=6,max=25"`
	ConfirmPassword string     `json:"confirm_password" validate:"required,eqfield=Password"`
	Gender          *bool      `json:"gender"`
	FirstName       string     `json:"first_name" validate:"required"`
	LastName        *string    `json:"last_name"`
	DateOfBirth     *time.Time `json:"date_of_birth"`
}

type Login struct {
	ClientSecret string `json:"client_secret" validate:"required"`

	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Scope    string `json:"scope"`
}

type UpdatePassword struct {
	ID                      string `json:"id"`
	OldPassword             string `json:"old_password" validate:"required"`
	NewPassword             string `json:"new_password" validate:"required,min=6,max=25"`
	NewPasswordConfirmation string `json:"new_password_confirmation" validate:"required,eqfield=NewPassword"`
}
