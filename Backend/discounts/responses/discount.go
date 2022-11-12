package responses

type Discount struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type Discounts struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Success *bool                    `json:"success"`
	Data    []map[string]interface{} `json:"data"`
}
