package responses

import "sales-service/models"

type Sale struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type Sales struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    []map[string]interface{} `json:"data"`
}

type SaleData struct {
	Success *bool                 `json:"success"`
	Data    []models.SaleListItem `json:"data"`
	Message string                `json:"message"`
}

type SaleInfo struct {
	Success *bool               `json:"success"`
	Data    models.SaleListItem `json:"data"`
	Message string              `json:"message"`
}
