package models

type MyGame struct {
	Username string `json:"username,omitempty" validate:"required"`
	IDGame   int64  `json:"idGame,omitempty" validate:"required"`
}

type MyGames struct {
	MyGamesList []MyGame `json:"mygames"`
}

type MyGamesUpdate struct {
	IDMyGame   int64 `json:"idMyGame,omitempty" validate:"required"`
	IDUser     int64 `json:"username,omitempty" validate:"required"`
	IDGame     int64 `json:"idGame,omitempty" validate:"required"`
	IsDeleted  *bool `json:"isDeleted,omitempty" validate:"omitempty"`
	IsWishlist int64 `json:"isWishlist,omitempty" validate:"required"`
	IsLibrary  int64 `json:"isLibrary,omitempty" validate:"required"`
}
