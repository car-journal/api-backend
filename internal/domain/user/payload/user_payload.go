package userpayload

type CreatePayload struct {
	Email    string `json:"email" validate:"required,email,max=255,email"`
	Password string `json:"password"`
}
