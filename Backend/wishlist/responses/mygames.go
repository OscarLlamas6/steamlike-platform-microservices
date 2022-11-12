package responses

type MyGame struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type MyGames struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    []map[string]interface{} `json:"data"`
}
