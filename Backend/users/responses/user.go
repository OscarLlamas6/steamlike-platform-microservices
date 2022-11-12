package responses

import "users-service/models"

type User struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type Users struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    []map[string]interface{} `json:"data"`
}

type UsersList struct {
	Success *bool                 `json:"success"`
	Data    []models.UserComplete `json:"data"`
	Message string                `json:"message"`
}
