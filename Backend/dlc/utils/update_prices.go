package utils

import (
	"dlc-service/configs"
	"fmt"
)

func UpdatePrices(prices []interface{}, idDLC string) bool {

	db := configs.ConnectDB()
	defer db.Close()

	_, err := db.Exec("UPDATE `RegionPrice` SET isDeleted = 1 WHERE idDLC = ?;", idDLC)
	if err != nil {
		defer db.Close()
		fmt.Printf("Error: %v\n", err)
		return false
	}

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

		defer myQuery2.Close()

		if myQuery2.Next() {
			err = myQuery2.Scan(&idRegion)
			if err != nil {
				defer db.Close()
				fmt.Printf("Error: %v\n", err)
				return false
			}
		}

		_, err = db.Exec("INSERT INTO RegionPrice (idDLC, idRegion,isDeleted, price, discount, isDLC) VALUES(?,?,?,?,?,?)", idDLC, idRegion, 0, price, discount, 1)
		if err != nil {
			defer db.Close()
			fmt.Printf("Error: %v\n", err)
			return false
		}
	}

	return true
}
