package userdto

import (
	internalmodel "github.com/car-journal/internal/model"
)

type UserWithUserProfile struct {
	internalmodel.User

	UserProfile *internalmodel.UserProfile `gorm:"->;foreignKey:user_id" json:"user_profile"`
}

func (u UserWithUserProfile) ToModel() *internalmodel.User {
	return &internalmodel.User{
		BaseModel: internalmodel.BaseModel{
			ID:              u.ID,
			CreatedAt:       u.CreatedAt,
			UpdatedAt:       u.UpdatedAt,
			DeletedAt:       u.DeletedAt,
			AffectedRecords: u.AffectedRecords,
		},
		Email:    u.Email,
		Password: u.Password,
	}
}
