package controllers

import (
	"apartment_rent/configs"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunRoutes() error {
	router := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", PingPong)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	announcementD := router.Group("/announcements/delete",checkUserAuthentication)
	{
		announcementD.DELETE("/:id",DeleteAnnouncement)
	}
	announcementR := router.Group("/announcements/review",checkUserAuthentication)
	{
		announcementR.DELETE("/:id",DeleteReview)
	}

	userG := router.Group("/users",checkUserAuthentication)
	{
		userG.GET("", GetAllUsers)
		userG.GET("/:id", GetUserByID)
		userG.POST("", CreateUser) 
		userG.PUT("/:id", UpdateUser)
		userG.DELETE("/:id",DeleteUser)
		userG.POST("/block/:id", BlockUser)
		
	}
	announcementG := router.Group("/announcements",checkUserAuthentication1)
	{
		announcementG.GET("", GetAllAnnouncement)
		announcementG.GET("/price/:price", GetAnnouncementByPrice)
		announcementG.POST("", CreateAnnouncement)
		announcementG.PUT("/:id", UpdateAnnouncement)
		announcementG.GET(":id", GetAnnouncementByID)
		announcementG.GET("/count_apart/:count_apart", GetAnnouncementByRooms)
		announcementG.GET("/views/report/:id", GetViewsReport)
		announcementG.POST("/review", CreateReview)
		announcementG.GET("/reviews/:id", GetReviews)
		announcementG.POST("/favorite/:id", AddToFavorites)
		announcementG.GET("/favorites", GetFavorites)
		
	}



	err := router.Run(fmt.Sprintf("%s:%s", configs.AppSettings.AppParams.ServerURL, configs.AppSettings.AppParams.PortRun))

	if err != nil {
		return err
	}

	return nil

}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

}
