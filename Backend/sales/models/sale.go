package models

type Sale struct {
	IDUser       int64         `json:"idUser,omitempty" validate:"required"`
	SaleDate     string        `json:"saleDate,omitempty" validate:"required"`
	Total        float64       `json:"total,omitempty" validate:"required"`
	MetodoDePago string        `json:"metodoDePago,omitempty" validate:"required"`
	Detalle      []interface{} `json:"detalle,omitempty" validate:"required"`
}

type Sales struct {
	SalesList []Sale `json:"sales"`
}

type SaleUpdate struct {
	IDSale       int64         `json:"idSale,omitempty" validate:"required"`
	IDUser       int64         `json:"idUser,omitempty" validate:"required"`
	SaleDate     string        `json:"saleDate,omitempty" validate:"required"`
	Total        float64       `json:"total,omitempty" validate:"required"`
	MetodoDePago string        `json:"metodoDePago,omitempty" validate:"required"`
	Detalle      []interface{} `json:"detalle,omitempty" validate:"omitempty"`
}
type SaleListItem struct {
	IDSale       int64        `json:"idSale,omitempty" validate:"required"`
	IDUser       int64        `json:"idUser,omitempty" validate:"required"`
	SaleDate     string       `json:"saleDate,omitempty" validate:"required"`
	Total        float64      `json:"total,omitempty" validate:"required"`
	MetodoDePago string       `json:"metodoDePago,omitempty" validate:"required"`
	Detalle      []SaleDetail `json:"detalle,omitempty" validate:"required"`
}

type SaleDetail struct {
	IDDetail int64   `json:"idDetail,omitempty" validate:"required"`
	IDSale   int64   `json:"idSale,omitempty" validate:"required"`
	IdGame   int64   `json:"idGame,omitempty" validate:"omitempty"`
	IdDLC    int64   `json:"idDLC,omitempty" validate:"omitempty"`
	SubTotal float64 `json:"subTotal,omitempty" validate:"required"`
	IsDLC    int64   `json:"isDLC,omitempty" validate:"required"`
}
