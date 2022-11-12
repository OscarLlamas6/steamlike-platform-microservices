package models

type SaleMail struct {
	Email    string        `json:"email" validate:"required"`
	UserName string        `json:"userName" validate:"required"`
	SaleDate string        `json:"saleDate,omitempty" validate:"required"`
	Total    float64       `json:"price" validate:"required"`
	Details  []interface{} `json:"details,omitempty" validate:"omitempty"`
}
