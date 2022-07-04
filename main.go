package main

import (
	"example/shop-golang/model"
	"example/shop-golang/middleware/jwt"
	"example/shop-golang/router"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// connect to database
	fmt.Println("Connect to database")
	db, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Migrate database")
	db.AutoMigrate(&model.Production{})
	fmt.Println("Save to database")

	// setup router
	r := gin.Default()

	// middleware
	r.Use(jwt.JwtMiddleware)

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
		productionRouter.GET("", router.GetAllProduction)
		productionRouter.POST("", router.CreateProduction)
	}

	// start server
	r.Run()
}
