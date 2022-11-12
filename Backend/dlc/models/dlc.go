package models

type DLC struct {
	Name           string        `json:"name,omitempty" validate:"required"`
	IDGame         int64         `json:"idGame,omitempty" validate:"required"`
	ImageURL       string        `json:"imageURL,omitempty" validate:"omitempty"`
	ReleaseDate    string        `json:"releaseDate,omitempty" validate:"required"`
	Description    string        `json:"description,omitempty" validate:"required"`
	IsGlobal       int64         `json:"isGlobal,omitempty" validate:"gte=0"`
	GlobalPrice    float64       `json:"globalPrice,omitempty" validate:"gte=0"`
	GlobalDiscount float64       `json:"globalDiscount,omitempty" validate:"gte=0"`
	Prices         []interface{} `json:"prices,omitempty" validate:"omitempty"`
}

type DLCs struct {
	DLCList []DLC `json:"dlcs"`
}

type DLCUpdate struct {
	IDDLC          int64         `json:"idDLC,omitempty" validate:"required"`
	Name           string        `json:"name,omitempty" validate:"required"`
	IDGame         int64         `json:"idGame,omitempty" validate:"omitempty"`
	ImageURL       string        `json:"imageURL,omitempty" validate:"omitempty"`
	ReleaseDate    string        `json:"releaseDate,omitempty" validate:"required"`
	Description    string        `json:"description,omitempty" validate:"required"`
	IsGlobal       int64         `json:"isGlobal,omitempty" validate:"gte=0"`
	GlobalPrice    float64       `json:"globalPrice,omitempty" validate:"gte=0"`
	GlobalDiscount float64       `json:"globalDiscount,omitempty" validate:"gte=0"`
	Prices         []interface{} `json:"prices,omitempty" validate:"omitempty"`
	UpdateImage    int64         `json:"updateImage,omitempty" validate:"omitempty"`
}
type DLCListItem struct {
	IDDLC            int64                `json:"idDLC,omitempty" validate:"required"`
	Name             string               `json:"name,omitempty" validate:"required"`
	IDGame           int64                `json:"idGame,omitempty" validate:"required"`
	ImageURL         string               `json:"imageURL,omitempty" validate:"omitempty"`
	ReleaseDate      string               `json:"releaseDate,omitempty" validate:"required"`
	Description      string               `json:"description,omitempty" validate:"required"`
	IsGlobal         *bool                `json:"isGlobal,omitempty" validate:"required"`
	IsDeleted        *bool                `json:"isDeleted,omitempty" validate:"omitempty"`
	GlobalPrice      float64              `json:"globalPrice,omitempty" validate:"gte=0"`
	GlobalDiscount   float64              `json:"globalDiscount,omitempty" validate:"gte=0"`
	UpdateImage      int64                `json:"updateImage,omitempty" validate:"omitempty"`
	RegionsAndPrices []RegionGameListItem `json:"region,omitempty" validate:"required"`
}

type RegionGameListItem struct {
	IDRegion int64   `json:"idRegion,omitempty" validate:"required"`
	Region   string  `json:"region,omitempty" validate:"required"`
	Discount float64 `json:"discount,omitempty" validate:"required"`
	Price    float64 `json:"price,omitempty" validate:"required"`
}
