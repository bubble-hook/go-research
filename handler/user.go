package handler

import (
	"go-x/model"

	"github.com/gin-gonic/gin"
)

// CreateUser : CreateUser
func (x Handler) CreateUser(c *gin.Context) {
	createUserRequest := new(model.User)

	// binding request body to createUserRequest
	if err := c.ShouldBind(createUserRequest); err != nil {
		c.JSON(400, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	//
	if err := x.db.Create(createUserRequest).Error; err != nil {
		c.JSON(500, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, map[string]interface{}{
		"status": "success",
		"result": createUserRequest,
	})

}

// GetUsers : GetUsers
func (h Handler) GetUsers(c *gin.Context) {

	users := make([]model.User, 0)

	if err := h.db.Find(&users).Error; err != nil {
		c.JSON(500, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	reponseUsers := make([]model.UserReponse, 0)

	for _, user := range users {
		reponseUsers = append(reponseUsers, *user.ParseToReponse())
	}

	c.JSON(200, map[string]interface{}{
		"status": "success",
		"result": reponseUsers,
	})
}

// GetUserByID : GetUserByID
func (h Handler) GetUserByID(c *gin.Context) {

	c.JSON(200, map[string]interface{}{
		"test": "test",
	})
}
