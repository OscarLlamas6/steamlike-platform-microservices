package utils

import (
	"fmt"
	"games-service/configs"
)

func SavePrices(prices []interface{}, idGame string) bool {

	db := configs.ConnectDB()

	for _, price := range prices {

		priceObj := price.(map[string]interface{})
		region := priceObj["region"].(string)
		price := priceObj["price"].(float64)
		discount := priceObj["discount"].(float64)

		///////// GETTING REGION ID

		myQuery2, err := db.Query("SELECT idRegion FROM Region WHERE name = ?;", region)
		if err != nil {
			defer db.Close()
			fmt.Printf("Error: %v\n", err)
			return false
		}

		var idRegion int64 = 1

		if myQuery2.Next() {

			err = myQuery2.Scan(&idRegion)
			if err != nil {
				defer db.Close()
				fmt.Printf("Error: %v\n", err)
				return false
			}
		}

		myQuery2.Close()

		_, err = db.Exec("INSERT INTO RegionPrice (idGame, idRegion, isDeleted, price, discount, isDLC) VALUES(?,?,?,?,?,?)", idGame, idRegion, 0, price, discount, 0)
		if err != nil {
			defer db.Close()
			fmt.Printf("Error: %v\n", err)
			return false
		}
	}
	defer db.Close()
	return true
}
