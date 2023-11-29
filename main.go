package main

import (
	"qq-doctor-message/routes"
	"qq-doctor-message/utils"
)

func main() {
	utils.InitMysql()
	utils.InitConfig()

	r := routes.Router()
	err := r.Run(":8081")
	if err != nil {
		return
	}
}
