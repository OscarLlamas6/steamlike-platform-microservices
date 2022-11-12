package utils

import (
	"fmt"
	"games-service/configs"
)

func SaveDevelopers(developers []interface{}, idGame string) bool {

	db := configs.ConnectDB()
	defer db.Close()

	for _, id := range developers {
		idDeveloper := int64(id.(float64))
		_, err := db.Exec("INSERT INTO GameDeveloper (idGame, idDeveloper, isDeleted) VALUES(?,?,?)", idGame, idDeveloper, 0)
		if err != nil {
			defer db.Close()
			fmt.Printf("Error: %v\n", err)
			return false
		}
	}

	return true
}
