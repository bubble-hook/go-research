package main

import (
	"go-x/db"
	"go-x/handler"
	"go-x/model"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := db.GetDbConnect() // db connect and error
	if err != nil {
		panic(err.Error()) // the program stop
	}

	db.AutoMigrate(&model.User{}) //

	gin := gin.New()

	handler.Init(db, gin)

	gin.Run(":9898")

}
