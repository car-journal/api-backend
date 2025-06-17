package internalmodel

import (
	"time"

	"github.com/car-journal/lib/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID              uuid.UUID      `json:"id" gorm:"primarykey"`
	CreatedAt       time.Time      `json:"created_at" gorm:"<-:created"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" sql:"index"`
	AffectedRecords int64          `json:"-" gorm:"<-:false"` // count affected recrods that satisfied the DB query (exclude offset and limit)
}

type BaseModelNoSoftDelete struct {
	ID              uuid.UUID `json:"id" gorm:"primarykey"`
	CreatedAt       time.Time `json:"created_at" gorm:"<-:created"`
	UpdatedAt       time.Time `json:"updated_at"`
	AffectedRecords int64     `json:"-" gorm:"<-:false"` // count affected recrods that satisfied the DB query (exclude offset and limit)
}

type BaseModelNoID struct {
	CreatedAt       time.Time      `json:"created_at" gorm:"<-:created"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" sql:"index"`
	AffectedRecords int64          `json:"-" gorm:"<-:false"` // count affected recrods that satisfied the DB query (exclude offset and limit)
}
