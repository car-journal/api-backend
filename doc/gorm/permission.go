package gorm

type Permission struct {
	_ string `gorm:"<-:create"`          // allow read and create
	_ string `gorm:"<-:update"`          // allow read and update
	_ string `gorm:"<-"`                 // allow read and write (create and update)
	_ string `gorm:"<-:false"`           // allow read, disable write permission
	_ string `gorm:"->"`                 // readonly (disable write permission unless it configured)
	_ string `gorm:"->;<-:create"`       // allow read and create
	_ string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
	_ string `gorm:"-"`                  // ignore this field when write and read with struct
	_ string `gorm:"-:all"`              // ignore this field when write, read and migrate with struct
	_ string `gorm:"-:migration"`        // ignore this field when migrate with struct
}
