package models

import (
	"fmt"
	"gorm.io/gorm"
	"qq-doctor-message/utils"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time `gorm:"type:datetime(0)"`
	HeartbeatTime time.Time `gorm:"type:datetime(0)"`
	LoginOutTime  time.Time `gorm:"type:datetime(0)"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Print(v)
	}
	return data
}
