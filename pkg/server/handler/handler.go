package handler

import "github.com/gin-gonic/gin"

func HandleLogCreate() gin.HandlerFunc{
	return func(c *gin.Context){
		// ここに処理を書いていく
		c.JSON(200, gin.H{
			"log_id":"asdfghjkl",
		})
	}
}