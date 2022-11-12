package responses

type Developer struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type Developers struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Success *bool                    `json:"success"`
	Data    []map[string]interface{} `json:"data"`
}
