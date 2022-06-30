package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Production struct {
	gorm.Model
	Name   string `json:"name"`
	Price  uint   `json:"price"`
	Amount uint   `json:"amount"`
}

func main() {

	// connect to database
	fmt.Println("Connect to database")
	db, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Migrate database")
	db.AutoMigrate(&Production{})
	fmt.Println("Save to database")

	// setup router
	r := gin.Default()

	// home router
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name":    "Shop",
			"version": "0.0.1",
		})
	})

	// productions router
	productionRouter := r.Group("/productions")
	{
		productionRouter.GET("", GetAllProduction)
		productionRouter.POST("", CreateProduction)
	}

	// start server
	r.Run()
}

func ConnectToDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}

func GetAllProduction(c *gin.Context) {
	db := ConnectToDatabase()
	var productions []Production
	db.Find(&productions)
	c.IndentedJSON(http.StatusOK, productions)
}

func CreateProduction(c *gin.Context) {
	db := ConnectToDatabase()
	production := &Production{}
	err := c.ShouldBindJSON(production)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	db.Create(production)
	c.AbortWithStatus(http.StatusNoContent)
}
