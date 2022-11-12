package models

type Game struct {
	Name           string        `json:"name,omitempty" validate:"required"`
	ImageURL       string        `json:"imageURL,omitempty" validate:"omitempty"`
	ReleaseDate    string        `json:"releaseDate,omitempty" validate:"required"`
	RestrictionAge string        `json:"restrictionAge,omitempty" validate:"required"`
	Description    string        `json:"description,omitempty" validate:"required"`
	Group          int64         `json:"group,omitempty" validate:"gte=0"`
	Prices         []interface{} `json:"prices,omitempty" validate:"required"`
	IsGlobal       int64         `json:"isGlobal,omitempty" validate:"gte=0"`
	GlobalPrice    float64       `json:"globalPrice,omitempty" validate:"gte=0"`
	GlobalDiscount float64       `json:"globalDiscount,omitempty" validate:"gte=0"`
	Developers     []interface{} `json:"developers,omitempty" validate:"required"`
	Categories     []interface{} `json:"categories,omitempty" validate:"required"`
}

type Games struct {
	GamesList []Game `json:"games"`
}

type GamesUpdate struct {
	IDGame         int64         `json:"idGame,omitempty" validate:"gte=1"`
	Name           string        `json:"name,omitempty" validate:"required"`
	ImageURL       string        `json:"imageURL,omitempty" validate:"omitempty"`
	ReleaseDate    string        `json:"releaseDate,omitempty" validate:"required"`
	RestrictionAge string        `json:"restrictionAge,omitempty" validate:"required"`
	Description    string        `json:"description,omitempty" validate:"required"`
	Group          int64         `json:"group,omitempty" validate:"required"`
	IsDeleted      *bool         `json:"isDeleted,omitempty" validate:"omitempty"`
	Prices         []interface{} `json:"prices,omitempty" validate:"omitempty"`
	IsGlobal       int64         `json:"isGlobal,omitempty" validate:"gte=0"`
	GlobalPrice    float64       `json:"globalPrice,omitempty" validate:"gte=0"`
	GlobalDiscount float64       `json:"globalDiscount,omitempty" validate:"gte=0"`
	Developers     []interface{} `json:"developers,omitempty" validate:"omitempty"`
	Categories     []interface{} `json:"categories,omitempty" validate:"omitempty"`
	UpdateImage    int64         `json:"updateImage,omitempty" validate:"omitempty"`
}
type GameListItem struct {
	GameID         int64                   `json:"gameId,omitempty" validate:"gte=1"`
	Name           string                  `json:"name,omitempty" validate:"required"`
	Image          string                  `json:"image,omitempty" validate:"required"`
	Description    string                  `json:"description,omitempty" validate:"required"`
	ReleaseDate    string                  `json:"releaseDate,omitempty" validate:"required"`
	RestrictionAge string                  `json:"restrictionAge,omitempty" validate:"required"`
	Group          int64                   `json:"group,omitempty" validate:"required"`
	Discount       float64                 `json:"discount,omitempty" validate:"gte=0"`
	Price          float64                 `json:"price,omitempty" validate:"gte=0"`
	IsGlobal       *bool                   `json:"isGlobal,omitempty" validate:"required"`
	IsDeleted      *bool                   `json:"isDeleted,omitempty" validate:"omitempty"`
	Developer      []DeveloperGameListItem `json:"developer,omitempty" validate:"required"`
	Category       []string                `json:"category,omitempty" validate:"required"`
	Regions        []string                `json:"regions,omitempty" validate:"required"`
	Prices         []RegionGameListItem    `json:"prices,omitempty" validate:"required"`
}

type FullGameListItem struct {
	GameID           int64                   `json:"gameId,omitempty" validate:"gte=1"`
	Name             string                  `json:"name,omitempty" validate:"required"`
	Image            string                  `json:"image,omitempty" validate:"required"`
	Description      string                  `json:"description,omitempty" validate:"required"`
	ReleaseDate      string                  `json:"releaseDate,omitempty" validate:"required"`
	RestrictionAge   string                  `json:"restrictionAge,omitempty" validate:"required"`
	Group            int64                   `json:"group,omitempty" validate:"required"`
	Discount         float64                 `json:"discount,omitempty" validate:"required"`
	Price            float64                 `json:"price,omitempty" validate:"required"`
	IsGlobal         *bool                   `json:"isGlobal,omitempty" validate:"required"`
	Developers       []DeveloperGameListItem `json:"developers,omitempty" validate:"required"`
	Categories       []CategoryItem          `json:"categories,omitempty" validate:"required"`
	RegionsAndPrices []RegionGameListItem    `json:"regions,omitempty" validate:"required"`
}

type CategoryItem struct {
	CategoryID int64  `json:"idCategory,omitempty" validate:"gte=1"`
	Name       string `json:"name,omitempty" validate:"required"`
}

type DeveloperGameListItem struct {
	DeveloperID int64  `json:"developerId,omitempty" validate:"gte=1"`
	Name        string `json:"name,omitempty" validate:"required"`
	Image       string `json:"image,omitempty" validate:"required"`
}

type RegionGameListItem struct {
	Region   string  `json:"region,omitempty" validate:"required"`
	Discount float64 `json:"discount,omitempty" validate:"gte=0"`
	Price    float64 `json:"price,omitempty" validate:"gte=0"`
}
