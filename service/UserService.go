package service

import (
	"github.com/gin-gonic/gin"
	"qq-doctor-message/models"
)

func GetUserList(ctx *gin.Context) {
	data := models.GetUserList()
	ctx.JSON(200,
		gin.H{
			"message": data,
		})

}
