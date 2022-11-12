package utils

import (
	"fmt"
	"games-service/configs"
)

func SaveCategories(categories []interface{}, idGame string) bool {

	db := configs.ConnectDB()
	defer db.Close()

	for _, id := range categories {
		idCategory := int64(id.(float64))
		_, err := db.Exec("INSERT INTO GameCategory (idGame, idCategory, isDeleted) VALUES(?,?,?)", idGame, idCategory, 0)
		if err != nil {
			defer db.Close()
			fmt.Printf("Error: %v\n", err)
			return false
		}
	}

	return true
}
