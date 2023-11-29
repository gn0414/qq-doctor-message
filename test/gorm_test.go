package test

import (
	"fmt"
	"qq-doctor-message/models"
	"qq-doctor-message/utils"
	"testing"
	"time"
)

func TestGorm(t *testing.T) {

	db := utils.DB

	db.AutoMigrate(&models.UserBasic{})

	user := &models.UserBasic{}
	user.Name = "熙蒙"

	user.LoginOutTime = time.Now()
	user.LoginTime = time.Now()
	user.HeartbeatTime = time.Now()
	db.Create(user)

	fmt.Print(db.First(user, 1))

	db.Model(user).Update("PassWord", "1234")

}
