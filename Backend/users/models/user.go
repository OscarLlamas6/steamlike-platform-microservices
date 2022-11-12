package models

type User struct {
	Name      string `json:"name,omitempty" validate:"required"`
	LastName  string `json:"lastName,omitempty" validate:"required"`
	UserName  string `json:"userName,omitempty" validate:"required"`
	BirthDate string `json:"birthDate,omitempty" validate:"required"`
	Email     string `json:"email,omitempty" validate:"email,required"`
	Password  string `json:"password,omitempty" validate:"required"`
	Region    string `json:"region,omitempty" validate:"required"`
}

type Users struct {
	Students []User `json:"students"`
}

type UserUpdate struct {
	Id          int64  `json:"id,omitempty" validate:"required"`
	UserName    string `json:"userName,omitempty" validate:"required"`
	Name        string `json:"name,omitempty" validate:"required"`
	LastName    string `json:"lastName,omitempty" validate:"required"`
	BirthDate   string `json:"birthDate,omitempty" validate:"required"`
	Email       string `json:"email,omitempty" validate:"email,required"`
	Region      int64  `json:"region,omitempty" validate:"required"`
	OldPassword string `json:"oldPassword,omitempty" validate:"required"`
	NewPassword string `json:"newPassword,omitempty" validate:"omitempty"`
	SetNewPass  int64  `json:"setNewPass,omitempty" validate:"omitempty"`
}

type UserPayload struct {
	Id       int64  `json:"id,omitempty" validate:"required"`
	Name     string `json:"name,omitempty" validate:"required"`
	LastName string `json:"lastName,omitempty" validate:"required"`
	UserName string `json:"userName,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"email,required"`
	Region   string `json:"region,omitempty" validate:"required"`
}

type UserComplete struct {
	Id          int64  `json:"id,omitempty" validate:"required"`
	Name        string `json:"name,omitempty" validate:"required"`
	LastName    string `json:"lastName,omitempty" validate:"required"`
	UserName    string `json:"userName,omitempty" validate:"required"`
	BirthDate   string `json:"birthDate,omitempty" validate:"required"`
	Email       string `json:"email,omitempty" validate:"email,required"`
	VerifyToken string `json:"verifyToken,omitempty" validate:"required"`
	Password    string `json:"password,omitempty" validate:"required"`
	IsActive    *bool  `json:"isActive,omitempty" validate:"required"`
	TimeOut     int64  `json:"timeOut,omitempty" validate:"required"`
	IsDeleted   *bool  `json:"isDeleted,omitempty" validate:"required"`
	ImageURL    string `json:"imageUrl,omitempty" validate:"required"`
	Region      int64  `json:"region,omitempty" validate:"required"`
}
