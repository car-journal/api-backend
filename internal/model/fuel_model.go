package internalmodel

type Fuel struct {
	BaseModel

	Brand string  `json:"brand"`
	Type  string  `json:"type"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
