package gorm

type Association struct {
	_ string `gorm:"foreignKey:value"`       // Specifies column name of the current model that is used as a foreign key to the join table
	_ string `gorm:"references:value"`       // Specifies column name of the reference’s table that is mapped to the foreign key of the join table
	_ string `gorm:"polymorphic:value"`      // Specifies polymorphic type such as model name
	_ string `gorm:"polymorphicValue:value"` // Specifies polymorphic value, default table name
	_ string `gorm:"many2many:value"`        // Specifies join table name
	_ string `gorm:"joinForeignKey:value"`   // Specifies foreign key column name of join table that maps to the current table
	_ string `gorm:"joinReferences:value"`   // Specifies foreign key column name of join table that maps to the reference’s table
	_ string `gorm:"constraint:value"`       // Relations constraint, e.g: OnUpdate,OnDelete
}
