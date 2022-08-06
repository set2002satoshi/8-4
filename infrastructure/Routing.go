package infrastructure

import (
	"github.com/gin-gonic/gin"

	"github.com/set2002satoshi/8-4/interfaces/controllers"
)

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *DB) *Routing {
	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: ":8080",
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
    usersController := controllers.NewUsersController(r.DB)
    r.Gin.GET("/users/:id", func (c *gin.Context) { usersController.Get(c) })
	r.Gin.POST("/users", func (c *gin.Context) { usersController.Post(c) })
}


