package infrastructure

import (
	"github.com/gin-gonic/gin"

	"github.com/set2002satoshi/8-4/interfaces/controllers/user"
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
	usersController := user.NewUsersController(r.DB)
	r.Gin.GET("/api/users/:id", func(c *gin.Context) { usersController.FindByID(c) })
	r.Gin.POST("/api/users", func(c *gin.Context) { usersController.Create(c) })
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
