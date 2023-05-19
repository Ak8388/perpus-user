package main

import (
	"latihangolanguser/connection"
	"latihangolanguser/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Db := connection.Connection()
	eng := routes.Routers{
		R:  r,
		DB: Db,
	}

	eng.Routs()
	r.Run()
}
