package models

type Developer struct {
	Name     string `json:"name,omitempty" validate:"required"`
	Pais     string `json:"pais,omitempty" validate:"required"`
	ImageURL string `json:"image,omitempty" validate:"omitempty"`
	Email    string `json:"email,omitempty" validate:"required"`
}

type Developers struct {
	DevelopersList []Developer `json:"developers"`
}

type DeveloperComplete struct {
	IdDeveloper int64  `json:"idDeveloper,omitempty" validate:"required"`
	Name        string `json:"name,omitempty" validate:"required"`
	Pais        string `json:"pais,omitempty" validate:"required"`
	ImageURL    string `json:"image,omitempty" validate:"omitempty"`
	Email       string `json:"email,omitempty" validate:"required"`
	IsDeleted   *bool  `json:"isDeleted,omitempty" validate:"omitempty"`
	UpdateImage int64  `json:"updateImage,omitempty" validate:"omitempty"`
}
