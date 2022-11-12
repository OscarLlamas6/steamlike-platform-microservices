package utils

import (
	"sales-service/configs"
)

func SaveDetails(details []interface{}, idSale string) bool {

	db := configs.ConnectDB()
	defer db.Close()

	for _, detail := range details {

		detailsObj := detail.(map[string]interface{})
		isDLC := detailsObj["isDLC"].(int64)

		if isDLC == 0 {

			idGame := detailsObj["idGame"].(int64)
			subTotal := detailsObj["subTotal"].(float64)

			////// REGISTRAR NUEVO DETALLE DE JUEGO VENDIDO
			myQuery, err := db.Prepare("INSERT INTO SaleDetail (idSale, idGame, subTotal, isDeleted, isDLC) VALUES(?,?,?,?,?)")
			if err != nil {
				defer db.Close()
				return false
			}
			_, err = myQuery.Exec(idSale, idGame, subTotal, 0, 0)
			if err != nil {
				defer db.Close()
				return false
			}

		} else {

			idDLC := detailsObj["idDLC"].(int64)
			subTotal := detailsObj["subTotal"].(float64)

			////// REGISTRAR NUEVO DETALLE DE JUEGO VENDIDO
			myQuery, err := db.Prepare("INSERT INTO SaleDetail (idSale, idDLC, subTotal, isDeleted, isDLC) VALUES(?,?,?,?,?)")
			if err != nil {
				defer db.Close()
				return false
			}
			_, err = myQuery.Exec(idSale, idDLC, subTotal, 0, 1)
			if err != nil {
				defer db.Close()
				return false
			}

		}

	}

	return true
}
