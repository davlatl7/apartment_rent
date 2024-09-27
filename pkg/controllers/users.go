package controllers

import (
	"apartment_rent/logger"
	"apartment_rent/models"
	"apartment_rent/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllUsers
// @Summary Get All Users
// @Security ApiKeyAuth
// @Tags users
// @Description get list of all users
// @ID get-all-users
// @Produce json
// @Param q query string false "fill if you need search"
// @Success 200 {array} models.User
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	logger.Info.Printf("Client with ip: [%s] requested list of users\n", c.ClientIP())
	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
	logger.Info.Printf("Client with ip: [%s] got list of users\n", c.ClientIP())
}

// GetUserByID
// @Summary Get User By ID
// @Security ApiKeyAuth
// @Tags users
// @Description get user by id
// @ID get-user-by-id
// @Produce json
// @Param id path integer true "number of id user"
// @Success 200 {array} models.User
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetUserByID] invalid user_id path parameter: %s\n", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	user, err := service.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)

}

// CreateUser
// @Summary Create User
// @Security ApiKeyAuth
// @Tags users
// @Description create new user
// @ID create-user
// @Accept json
// @Produce json
// @Param input body models.User true "new user info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
	})

}

// UpdateUser
// @Summary Update User
// @Security ApiKeyAuth
// @Tags users
// @Description update existed user
// @ID update-user
// @Accept json
// @Produce json
// @Param id path integer true "id of the users"
// @Param input body models.User true "user update info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid announcement id",
		})
		return
	}

	var userUpdate models.UserFilterForUpdate
	if err := c.BindJSON(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	updateduser, err := service.UpdateUser(id, userUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "announcement updated successfully",
		"announcement":  updateduser,
	})
}
// DeleteUser
// @Summary Delete User
// @Security ApiKeyAuth
// @Tags users
// @Description delete an existing user
// @ID delete-user
// @Param id path integer true "id of the user to delete"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}
	err = service.DeleteUser(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	
	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}
// BlockUser godoc
// @Summary Block a user
// @Description Blocks a specific user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path uint true "User ID"
// @Success 200 {object} map[string]string "message: user blocked successfully"
// @Failure 500 {object} map[string]string "error: failed to block user"
// @Router /users/{id}/block [post]
// @Security BearerAuth
func BlockUser(c *gin.Context) {
	userID := c.Param("id")

	err := service.BlockUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to block user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user blocked successfully",
	})
}