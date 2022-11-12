package models

// Categories

type Category struct {
	Name string `json:"name,omitempty" validate:"required"`
}

type Categories struct {
	CategoriesList []Category `json:"categories,omitempty" validate:"required"`
}

type CategoryUpdate struct {
	IDCategory int64  `json:"idCategory,omitempty" validate:"required"`
	Name       string `json:"name,omitempty" validate:"required"`
	IsDeleted  *bool  `json:"isDeleted,omitempty" validate:"omitempty"`
}

// Regions

type Region struct {
	Name string `json:"name,omitempty" validate:"required"`
}

type Regions struct {
	RegionsList []Region `json:"regions,omitempty" validate:"required"`
}

type RegionUpdate struct {
	IDRegion  int64  `json:"idRegion,omitempty" validate:"required"`
	Name      string `json:"name,omitempty" validate:"required"`
	IsDeleted *bool  `json:"isDeleted,omitempty" validate:"omitempty"`
}
