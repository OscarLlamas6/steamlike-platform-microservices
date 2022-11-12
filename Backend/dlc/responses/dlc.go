package responses

import "dlc-service/models"

type DLC struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type DLCList struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    []map[string]interface{} `json:"data"`
}

type DLCData struct {
	Success *bool                `json:"success"`
	Data    []models.DLCListItem `json:"data"`
	Message string               `json:"message"`
}

type DLCInfo struct {
	Success *bool              `json:"success"`
	Data    models.DLCListItem `json:"data"`
	Message string             `json:"message"`
}
