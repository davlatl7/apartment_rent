package controllers

import (
	"apartment_rent/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userIDCtx           = "userID"
	userRoleCtx         = "userRole"
	userIsDeletedCtx    = "userIsDeleted"
)

func checkUserAuthentication(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "empty auth header",
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid auth header",
		})
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token is empty",
		})
		return
	}

	accessToken := headerParts[1]

	claims, err := service.ParseToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(claims)

	// Проверка на заблокированного пользователя
	if claims.IsDeleted {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "user is deleted",
		})
		return
	}

	// Проверка роли пользователя
	if claims.Role != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "only admin can access this resource",
		})
		return
	}

	c.Set(userIDCtx, claims.UserID)
	c.Set(userRoleCtx, claims.Role)
	c.Set(userIsDeletedCtx, claims.IsDeleted)
	c.Next()
}
func checkUserAuthentication1(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "empty auth header",
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid auth header",
		})
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token is empty",
		})
		return
	}

	accessToken := headerParts[1]

	claims, err := service.ParseToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(claims)

	userID := claims.UserID
	c.Set(userIDCtx, userID)	
	c.Set(userIDCtx, claims.UserID)
	c.Set(userRoleCtx, claims.Role)
	c.Next()
}
