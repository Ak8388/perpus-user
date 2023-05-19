package handlers

import (
	"latihangolanguser/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandlersUser(usecase auth.UsecaseAuth, r *gin.RouterGroup) {
	eng := handUser{usecase}

	v2 := r.Group("user")
	v2.POST("regist", eng.Regist)
	v2.POST("login", eng.Login)
}

type handUser struct {
	usecaseUser auth.UsecaseAuth
}

func (hand handUser) Regist(ctx *gin.Context) {
	if err := hand.usecaseUser.Regist(ctx); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Regist Succes"})
}

func (hand handUser) Login(ctx *gin.Context) {
	if err := hand.usecaseUser.Login(ctx); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Login Succes"})
}
