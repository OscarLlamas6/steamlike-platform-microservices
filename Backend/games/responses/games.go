package responses

import "games-service/models"

type Game struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type Games struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    []map[string]interface{} `json:"data"`
}

type GamesList struct {
	Success *bool                 `json:"success"`
	Data    []models.GameListItem `json:"data"`
	Message string                `json:"message"`
}

type GameInfo struct {
	Success *bool                   `json:"success"`
	Data    models.FullGameListItem `json:"data"`
	Message string                  `json:"message"`
}
