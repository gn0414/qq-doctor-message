package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("config mysql", viper.Get("mysql"))

}

func InitConfig() {
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")))
}
