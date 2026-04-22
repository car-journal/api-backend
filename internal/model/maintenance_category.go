package internalmodel

const MaintenanceCategoryTableName = "maintenance_categories"

type MaintenanceCategory struct {
	BaseModel

	Name        string  `json:"name"`
	Description *string `json:"description"`
}
