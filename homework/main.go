package main

import (
	"homework/app/midwares"
	"homework/config/database"
	"homework/config/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	router.Init(r)
	err:=r.Run()

	if err !=nil{
		log.Fatal("Serve start error:",err)
	}
}