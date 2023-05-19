package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routers struct {
	R  *gin.Engine
	DB *gorm.DB
}

func (r Routers) Routs() {
	v1 := r.R.Group("login-perpus-user")
}
