package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler Handler
type Handler struct {
	db *gorm.DB
}

// Init Init
func Init(db *gorm.DB, gin *gin.Engine) {

	// new Handler and assign db  to handler
	h := Handler{
		db: db,
	}

	// define router endpoint by struc func
	v1 := gin.Group("/v1")
	{
		v1.POST("/users", h.CreateUser)
		v1.GET("/users", h.GetUsers)
		v1.GET("/users/:id", h.GetUserByID)
	}

}

func CalCommission() {

}
