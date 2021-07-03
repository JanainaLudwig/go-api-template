package router

import (
	"go-api-template/config"
	"log"
	"net/http"
)

func StartApi() {
	for  {
		log.Println(http.ListenAndServe(":" + config.App.ApiPort, Routes()))
	}
}
