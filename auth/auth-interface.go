package auth

import (
	"latihangolanguser/entity"

	"github.com/gin-gonic/gin"
)

type RepoAuth interface {
	Regist(entity.User) error
	Login(string)
}

type UsecaseAuth interface {
	Regist(*gin.Context) error
	Login(*gin.Context) error
}
