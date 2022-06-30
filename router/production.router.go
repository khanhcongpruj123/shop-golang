package router

import (
	"example/shop-golang/model"
	productionRepo "example/shop-golang/repository/production"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProduction(c *gin.Context) {
	productions := productionRepo.GetAll()
	c.IndentedJSON(http.StatusOK, productions)
}

func CreateProduction(c *gin.Context) {
	production := &model.Production{}
	err := c.ShouldBindJSON(production)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	productionRepo.Create(production)
	c.AbortWithStatus(http.StatusNoContent)
}
