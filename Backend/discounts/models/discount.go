package models

type Discount struct {
	IdGame        int64   `json:"idGame,omitempty" validate:"omitempty"`
	IdDLC         int64   `json:"idDLC,omitempty" validate:"omitempty"`
	DiscountValue float64 `json:"discount,omitempty" validate:"required"`
	StartTime     string  `json:"startTime,omitempty" validate:"required"`
	EndTime       string  `json:"endTime,omitempty" validate:"required"`
	IsDLC         int64   `json:"isDLC,omitempty" validate:"required"`
}

type Discounts struct {
	DiscountsList []Discount `json:"discounts"`
}

type DiscountComplete struct {
	IdDiscount    int64   `json:"idDiscount,omitempty" validate:"required"`
	IdGame        int64   `json:"idGame,omitempty" validate:"omitempty"`
	IdDLC         int64   `json:"idDLC,omitempty" validate:"omitempty"`
	DiscountValue float64 `json:"discount,omitempty" validate:"required"`
	StartTime     string  `json:"startTime,omitempty" validate:"required"`
	EndTime       string  `json:"endTime,omitempty" validate:"required"`
	IsDLC         int64   `json:"isDLC,omitempty" validate:"required"`
	IsDeleted     *bool   `json:"isDeleted,omitempty" validate:"omitempty"`
}
