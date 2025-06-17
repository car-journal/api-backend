package internalmodel

const UserTableName = "users"

type User struct {
	BaseModel

	Email    string `json:"email"`
	Password string `json:"-"`
}
