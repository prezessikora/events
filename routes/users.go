package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prezessikora/events/models"
	"github.com/prezessikora/events/utils"
)

func createUser(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindBodyWithJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		fmt.Println(err)
		return
	}
	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not create user"})
		fmt.Println(err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "user created", "user": user})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindBodyWithJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		fmt.Println(err)
		return
	}
	err = user.VerifyCredetials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not generate JWT token"})
		fmt.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "login success", "token": token})

}
