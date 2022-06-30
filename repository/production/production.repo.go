package production

import (
	"example/shop-golang/database"
	"example/shop-golang/model"
)

func GetAll() (productions []model.Production) {
	db := database.ConnectToDatabase()
	db.Find(&productions)
	return
}

func Create(production *model.Production) {
	db := database.ConnectToDatabase()
	db.Create(production)
}
