package internalmodel

const AddressTableName = "addresses"

type Address struct {
	BaseModel

	Name              string  `json:"name"`
	AddressFirstLine  string  `json:"address_first_line"`
	AddressSecondLine *string `json:"address_second_line"`
	Country           *string `json:"country"`
	City              *string `json:"city"`
	Province          *string `json:"province"`
	District          *string `json:"district"`
	SubDistrict       *string `json:"sub_district"`
	NeighboorhoodUnit *int64  `json:"neighboorhood_unit"`
	CommunityUnit     *int64  `json:"community_unit"`
	PostalCode        *string `json:"postal_code"`
	Notes             *string `json:"notes"`
	IsMain            bool    `json:"is_main"`
}
