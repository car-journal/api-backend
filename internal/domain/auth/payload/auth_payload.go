package authpayload

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
