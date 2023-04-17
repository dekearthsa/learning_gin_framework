package controller

import (
	"learning_gin_framework/model"
	"github.com/gin-gonic/gin"
)

func TestPostMethod(c *gin.Context) {
	var res model.Member
	if err := c.BindJSON(&res); err != nil {
		return
	}
	member := model.Member{res.Name, res.Age, res.Phone}
	c.JSON(200, gin.H{"mem": member})
}
